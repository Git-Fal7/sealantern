package socket

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"net"
	"reflect"
	"sync"

	"github.com/git-fal7/sealantern/config"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/protocol/stream"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/pkg/component"

	"github.com/google/uuid"
)

type Conn struct {
	net.Conn
	Reader       *stream.ProtocolReader
	Username     string
	UUID         uuid.UUID
	Compression  bool
	KeepAlive    int
	PacketsQueue chan protocol.PacketOut

	ProxyData   []string
	packetMutex sync.Mutex

	State        types.State
	Protocol     int32
	Disconnected bool
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		Conn: conn,
		Reader: &stream.ProtocolReader{
			Reader: bufio.NewReader(conn),
		},
		Compression:  false,
		KeepAlive:    0,
		State:        types.HANDSHAKING,
		Disconnected: false,
	}
}

func (c *Conn) WritePacket(packetOut protocol.PacketOut) (err error) {
	if c.Disconnected {
		return
	}
	id := packet.GetPacketIDFromClientPacket(reflect.TypeOf(packetOut).Elem())
	if id == -1 {
		return
	}
	c.packetMutex.Lock()
	defer c.packetMutex.Unlock()
	return c.writePacket(packetOut, id)
}


func (c *Conn) writePacket(packet protocol.PacketOut, id int16) (err error) {
	if c.Compression {
		return c.writePacketWithCompression(packet, id)		
	}
	return c.writePacketWithoutCompression(packet, id)
}


/*
Without compression
Field Name 		Field Type 		Notes
Length 			VarInt 			Length of packet data + length of the packet ID
Packet ID 		VarInt 	
Data 			Byte Array 		Depends on the connection state and packet ID, see the sections below 
*/
func (c *Conn) writePacketWithoutCompression(packet protocol.PacketOut, id int16) (err error) {
	packetWriter := &stream.ProtocolWriter{}
	packetWriter.WriteVarInt(int(id)) // Packet ID
	packet.Write(packetWriter)		  // Data

	packetData := packetWriter.Bytes() // Packet ID + Data Uncompressed

	writer := &stream.ProtocolWriter{}
	writer.WriteVarInt(len(packetData)) // Length
	writer.WriteByteArray(packetData) 

	c.Conn.Write(writer.Bytes())

	if config.LanternConfig.Logs {
		if id != 0x26 {
			log.Printf("# <- %d %s %s", id, reflect.TypeOf(packet), fmt.Sprint(packet))
		}
	}
	return nil
}

/*
With Compression
Compressed? 	Field Name 		Field Type 		Notes
No 				Packet Length 	VarInt 			Length of Data Length + compressed length of (Packet ID + Data)
No 				Data Length 	VarInt 			Length of uncompressed (Packet ID + Data) or 0
Yes 			Packet ID 		Varint 			zlib compressed packet ID (see the sections below)
|__				Data 			Byte Array 		zlib compressed packet data (see the sections below)
*/

func (c *Conn) writePacketWithCompression(packet protocol.PacketOut, id int16) (err error) {
	packetWriter := &stream.ProtocolWriter{}
	packetWriter.WriteVarInt(int(id))
	packet.Write(packetWriter)

	packetData := packetWriter.Bytes()

	dataWriter := &stream.ProtocolWriter{}
	if len(packetData) < config.LanternConfig.Threshold {
		dataWriter.WriteVarInt(0) // Data Length
		dataWriter.Write(packetData)
	} else {
		var b bytes.Buffer
		var encoder *zlib.Writer
		if config.LanternConfig.CompressionLevel > 9 || config.LanternConfig.CompressionLevel < -1 {
			encoder, _ = zlib.NewWriterLevel(&b, -1)
		} else {
			encoder, _ = zlib.NewWriterLevel(&b, config.LanternConfig.CompressionLevel)
		}
		encoder.Write(packetData)
		encoder.Close()
		dataWriter.WriteVarInt(len(packetData))
		dataWriter.Write(b.Bytes())
	}

	writer := &stream.ProtocolWriter{}
	writer.WriteVarInt(dataWriter.Len())
	writer.Write(dataWriter.Bytes())
	c.Conn.Write(writer.Bytes())

	if config.LanternConfig.Logs {
		if id != 0x26 {
			log.Printf("# <- %d %s %s", id, reflect.TypeOf(packet), fmt.Sprint(packet))
		}
	}
	return nil
}

func (c *Conn) HandlePacket(id int, length int, reader *stream.ProtocolReader) (handledPacket protocol.PacketIn, err error) {
	typ := packet.GetPacketTypeFromRegistry(c.State, id)

	if typ == nil {
		if config.LanternConfig.Logs {
			log.Printf(" -> Unknown packet #%d\n", id)
		}

		var buff []byte
		nbr := 0
		if length > 500 {
			buff = make([]byte, 500)
		} else {
			buff = make([]byte, length)
		}

		for nbr < length {
			if length-nbr > 500 {
				reader.Read(buff)
				nbr += 500
			} else {
				reader.Read(buff[:length-nbr])
				nbr = length
			}
		}
		return nil, nil
	}

	handledPacket, _ = reflect.New(typ).Interface().(protocol.PacketIn)

	if err = handledPacket.Read(reader, length); err != nil {
		return nil, err
	}
	return
}

func (c *Conn) Disconnect(message component.IChatComponent) {
	if c.State == types.LOGIN {
		c.WritePacket(&packet.PacketLoginDisconnect{
			Component: message,
		})
	} else {
		c.WritePacket(&packet.PacketPlayDisconnect{
			Component: message,
		})
	}
	c.Disconnected = true
}

func (c *Conn) Active() bool {
	return !c.Disconnected
}
