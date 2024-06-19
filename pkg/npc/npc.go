package npc

import (
	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/world"
)

type NPC interface {
	EntityID() uint16
	Position() world.Position
	SendPackets(player.IPlayer)
	GetDestructionID() []uint16
}
