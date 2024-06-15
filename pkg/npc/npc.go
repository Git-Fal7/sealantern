package npc

import (
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/world"
)

type NPC interface {
	EntityID() uint16
	Position() world.Position
	GetCreationPacket() []protocol.PacketOut
	GetDestructionID() []uint16
}
