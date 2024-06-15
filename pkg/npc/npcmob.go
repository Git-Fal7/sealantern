package npc

import (
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
)

type NPCMob struct {
	entityID uint16
	mobType  types.MobType
	position world.Position
}

func NewNPCMob(entityID uint16, position world.Position, mobType types.MobType) *NPCMob {
	return &NPCMob{
		entityID: entityID,
		mobType:  mobType,
		position: position,
	}
}

func (npc NPCMob) EntityID() uint16 {
	return npc.entityID
}

func (npc NPCMob) Position() world.Position {
	return npc.position
}

func (npc NPCMob) GetCreationPacket() []protocol.PacketOut {
	return []protocol.PacketOut{
		&packet.PacketPlaySpawnMob{
			EntityID:  npc.entityID,
			MobType:   npc.mobType,
			Position:  npc.position,
			HeadPitch: npc.position.IntYaw(),
			VelocityX: 0,
			VelocityY: 0,
			VelocityZ: 0,
			Metadata:  make(metadata.MetadataMap),
		},
	}
}

func (npc NPCMob) GetDestructionID() []uint16 {
	return []uint16{
		npc.entityID,
	}
}
