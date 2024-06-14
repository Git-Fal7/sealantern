package npc

import (
	"crypto/rand"
	"fmt"

	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/google/uuid"
)

type NPCPlayer struct {
	entityID   uint16
	name       string
	uuid       uuid.UUID
	properties []profile.Property
	position   world.Position
}

func NewNPCPlayer(entityID uint16, position world.Position, properties []profile.Property) *NPCPlayer {
	b := make([]byte, 14)
	rand.Read(b)
	name := fmt.Sprintf("NPC-%x", b[2:6])
	npcUUID, _ := uuid.NewRandom()
	return &NPCPlayer{
		entityID:   entityID,
		name:       name,
		uuid:       npcUUID,
		properties: properties,
		position:   position,
	}
}

func (npc NPCPlayer) EntityID() uint16 {
	return npc.entityID
}

func (npc NPCPlayer) Position() world.Position {
	return npc.position
}

func (npc NPCPlayer) GetCreationPacket() []protocol.PacketOut {
	return []protocol.PacketOut{
		&packet.PacketPlayPlayerListItem{
			Action: types.PlayerListActionAddPlayer,
			Entries: []types.PlayerListEntry{
				{
					Profile: profile.PlayerProfile{
						UUID: npc.uuid,
						Name: npc.name,
					},
					GameMode:    types.SURVIVAL,
					Ping:        0,
					DisplayName: nil,
				},
			},
		},
		&packet.PacketPlaySpawnPlayer{
			EntityID:       npc.entityID,
			PlayerUUID:     npc.uuid,
			PlayerPosition: npc.position,
			CurrentItem:    0, // Air
		},
		&packet.PacketPlayEntityHeadLook{
			EntityID: npc.entityID,
			HeadYaw:  npc.position.IntYaw(),
		},
		&packet.PacketPlayTeams{
			TeamName:          npc.name,
			Mode:              types.TeamModeCreate,
			FriendlyFire:      types.TeamFriendlyFireOff,
			NameTagVisibility: types.TeamNameTagVisibilityNever,
			Color:             0,
			Players: []string{
				npc.name,
			},
		},
		&packet.PacketPlayPlayerListItem{
			Action: types.PlayerListActionRemovePlayer,
			Entries: []types.PlayerListEntry{
				{
					Profile: profile.PlayerProfile{
						UUID: npc.uuid,
						Name: npc.name,
					},
					GameMode:    types.SURVIVAL,
					Ping:        0,
					DisplayName: nil,
				},
			},
		},
	}
}
