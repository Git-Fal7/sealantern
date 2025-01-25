package npc

import (
	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
)

type NPCMob struct {
	entityID int32
	mobType  types.MobType
	position world.Position
}

func NewNPCMob(entityID int32, position world.Position, mobType types.MobType) *NPCMob {
	return &NPCMob{
		entityID: entityID,
		mobType:  mobType,
		position: position,
	}
}

func (npc NPCMob) EntityID() int32 {
	return npc.entityID
}

func (npc NPCMob) Position() world.Position {
	return npc.position
}

func (npc NPCMob) SendPackets(p player.IPlayer) {
	p.WritePacket(&packet.PacketPlaySpawnMob{
		EntityID:  npc.entityID,
		MobType:   npc.mobType,
		Position:  npc.position,
		HeadPitch: npc.position.IntYaw(),
		VelocityX: 0,
		VelocityY: 0,
		VelocityZ: 0,
		Metadata:  make(metadata.MetadataMap),
	})
}

func (npc NPCMob) GetDestructionID() []int32 {
	return []int32{
		npc.entityID,
	}
}
