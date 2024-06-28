package socket

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"reflect"

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

	ProxyData []string

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

func (c *Conn) WritePacket(packet protocol.PacketOut) (err error) {
	return c.writePacket(packet)
}

func (c *Conn) writePacket(packet protocol.PacketOut) (err error) {
	return c.writePacketWithoutCompression(packet)
}

func (c *Conn) writePacketWithoutCompression(packet protocol.PacketOut) (err error) {
	id := packet.Id()
	if id == -1 {
		return
	}
	packetWriter := &stream.ProtocolWriter{}
	packetWriter.WriteVarInt(int(id))
	packet.Write(packetWriter)

	packetData := packetWriter.Bytes()

	writer := &stream.ProtocolWriter{}
	writer.WriteVarInt(len(packetData))
	writer.WriteByteArray(packetData)

	c.Conn.Write(writer.Bytes())

	if config.LanternConfig.Logs {
		if packet.Id() != 0x26 {
			log.Printf("# <- %d %s %s", id, reflect.TypeOf(packet), fmt.Sprint(packet))
		}
	}
	return nil
}

func (c *Conn) HandlePacket(id int, length int) (handledPacket protocol.PacketIn, err error) {
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
				c.Reader.Read(buff)
				nbr += 500
			} else {
				c.Reader.Read(buff[:length-nbr])
				nbr = length
			}
		}
		return nil, nil
	}

	handledPacket, _ = reflect.New(typ).Interface().(protocol.PacketIn)

	if err = handledPacket.Read(c.Reader, length); err != nil {
		return nil, err
	}
	return
}

func (c *Conn) Disconnect(message component.IChatComponent) {
	msg, err := message.JSON()
	if err != nil {
		log.Fatal(err)
		return
	}
	if c.State == types.LOGIN {
		c.WritePacket(&packet.PacketLoginDisconnect{
			Component: msg,
		})
	} else {
		c.WritePacket(&packet.PacketPlayDisconnect{
			Component: msg,
		})
	}
	c.Disconnected = true
	c.Conn.Close()
}

func (c *Conn) Active() bool {
	return !c.Disconnected
}
