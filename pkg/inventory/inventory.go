package inventory

import "github.com/git-fal7/sealantern/minecraft/protocol"

type Inventory interface {
	ID() uint8
	Title() string
	Size() uint16
	Packets() []protocol.PacketOut
}
