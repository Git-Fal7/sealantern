package protocol

import (
	"github.com/git-fal7/sealantern/pkg/readerwriter"
)

type Packet interface {
}

// Server
type PacketIn interface {
	Packet
	Read(*readerwriter.ConnReadWrite, int) error
}

// Client
type PacketOut interface {
	Packet
	Id() int32
	Write(*readerwriter.ConnReadWrite) error
}
