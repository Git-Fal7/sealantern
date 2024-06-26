package protocol

import (
	"github.com/git-fal7/sealantern/minecraft/protocol/stream"
)

type Packet interface {
}

// Server
type PacketIn interface {
	Packet
	Read(*stream.ProtocolReader, int) error
}

// Client
type PacketOut interface {
	Packet
	Write(*stream.ProtocolWriter) error
}
