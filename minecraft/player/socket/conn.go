package socket

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/git-fal7/sealantern/config"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/readerwriter"

	"github.com/google/uuid"
)

type Conn struct {
	net.Conn
	IO *readerwriter.ConnReadWrite

	Username     string
	UUID         uuid.UUID
	Compression  bool
	KeepAlive    int
	PacketsQueue chan protocol.PacketOut

	State        types.State
	Protocol     int32
	Disconnected bool
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		Conn: conn,
		IO: &readerwriter.ConnReadWrite{
			Rdr: bufio.NewReader(conn),
			Wtr: bufio.NewWriter(conn),
		},
		Compression:  false,
		KeepAlive:    0,
		State:        types.HANDSHAKING,
		PacketsQueue: make(chan protocol.PacketOut),
		Disconnected: false,
	}
}

func (c *Conn) WritePacket(packet protocol.PacketOut) (err error) {
	c.PacketsQueue <- packet
	return nil
}

func (c *Conn) PrivateWritePacket(packet protocol.PacketOut) (err error) {
	return c.WritePacketWithoutCompression(packet)
}

func (c *Conn) WritePacketWithoutCompression(packet protocol.PacketOut) (err error) {
	buff := bytes.NewBuffer(nil) // 256
	tmp := c.IO
	c.IO = &readerwriter.ConnReadWrite{
		Rdr: tmp.Rdr,
		Wtr: buff,
	}

	id := packet.Id()
	if id == -1 {
		return
	}
	c.IO.WriteVarInt(int(id))
	packet.Write(c.IO)

	ln := bytes.NewBuffer(nil) // 256
	c.IO.Wtr = ln
	c.IO.WriteVarInt(buff.Len())
	c.IO = tmp
	c.Conn.Write(ln.Bytes())
	c.Conn.Write(buff.Bytes())

	if config.LanternConfig.Logs {
		log.Printf("# <- %d %s %s", id, reflect.TypeOf(packet), fmt.Sprint(packet))
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
				c.IO.Rdr.Read(buff)
				nbr += 500
			} else {
				c.IO.Rdr.Read(buff[:length-nbr])
				nbr = length
			}
		}
		return nil, nil
	}

	handledPacket, _ = reflect.New(typ).Interface().(protocol.PacketIn)

	if err = handledPacket.Read(c.IO, length); err != nil {
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
