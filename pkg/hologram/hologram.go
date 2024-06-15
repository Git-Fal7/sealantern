package hologram

import (
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
)

type Hologram struct {
	entityID    uint16
	DisplayName string
	position    world.Position
}

func NewHologram(entityID uint16, displayName string, position world.Position) *Hologram {
	return &Hologram{
		entityID:    entityID,
		DisplayName: displayName,
		position:    position,
	}
}

func (hologram Hologram) EntityID() uint16 {
	return hologram.entityID
}

func (hologram Hologram) Position() world.Position {
	return hologram.position
}

func (hologram Hologram) GetCreationPacket() []protocol.PacketOut {
	return []protocol.PacketOut{
		&packet.PacketPlaySpawnObject{
			EntityID:   hologram.entityID,
			ObjectType: types.ObjectTypeArmorStand,
			Position:   hologram.position,
			Data:       0,
		},
		&packet.PacketPlayEntityMetadata{
			EntityID: hologram.entityID,
			Metadata: metadata.MetadataMap{
				metadata.MetadataIndexStatus:     metadata.StatusFlagInvisible,
				metadata.MetadataNameTag:         hologram.DisplayName,
				metadata.MetadataShowNameTag:     uint8(1),
				metadata.MetadataArmorstandFlags: uint8(0x19), // Small armorstand, No baseplate, marker
			},
		},
	}
}
