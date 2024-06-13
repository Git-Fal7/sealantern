package protocol

import (
	"github.com/git-fal7/sealantern/pkg/readerwriter"
)

type Packet interface {
	Id() int32
}

// Server
type PacketIn interface {
	Packet
	Read(*readerwriter.ConnReadWrite, int) error
}

// Client
type PacketOut interface {
	Packet
	Write(*readerwriter.ConnReadWrite) error
}
