package npc

import (
	"crypto/rand"
	"fmt"

	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/google/uuid"
)

type NPC struct {
	EntityID uint16
	NPCType  NPCType
	Position world.Position
}

type NPCType types.MobType

const (
	NPCPlayer NPCType = 0
)

func (npc NPC) GetCreationPacket() []protocol.PacketOut {
	if npc.NPCType == NPCPlayer {
		b := make([]byte, 14)
		rand.Read(b)
		name := fmt.Sprintf("NPC-%x", b[2:6])
		npcUUID, _ := uuid.NewRandom()
		return []protocol.PacketOut{
			&packet.PacketPlayPlayerListItem{
				Action: types.PlayerListActionAddPlayer,
				Entries: []types.PlayerListEntry{
					{
						Profile: profile.PlayerProfile{
							UUID: npcUUID,
							Name: name,
						},
						GameMode:    types.SURVIVAL,
						Ping:        0,
						DisplayName: nil,
					},
				},
			},
			&packet.PacketPlaySpawnPlayer{
				EntityID:       npc.EntityID,
				PlayerUUID:     npcUUID,
				PlayerPosition: npc.Position,
				CurrentItem:    0, // Air
			},
			&packet.PacketPlayEntityHeadLook{
				EntityID: npc.EntityID,
				HeadYaw:  npc.Position.IntYaw(),
			},
			&packet.PacketPlayTeams{
				TeamName:          name,
				Mode:              types.TeamModeCreate,
				FriendlyFire:      types.TeamFriendlyFireOff,
				NameTagVisibility: types.TeamNameTagVisibilityNever,
				Color:             0,
				Players: []string{
					name,
				},
			},
			&packet.PacketPlayPlayerListItem{
				Action: types.PlayerListActionRemovePlayer,
				Entries: []types.PlayerListEntry{
					{
						Profile: profile.PlayerProfile{
							UUID: npcUUID,
							Name: name,
						},
						GameMode:    types.SURVIVAL,
						Ping:        0,
						DisplayName: nil,
					},
				},
			},
		}
	} else {
		return []protocol.PacketOut{
			&packet.PacketPlaySpawnMob{
				EntityID:  npc.EntityID,
				MobType:   types.MobType(npc.NPCType),
				Position:  npc.Position,
				HeadPitch: npc.Position.IntYaw(),
				VelocityX: 0,
				VelocityY: 0,
				VelocityZ: 0,
				Metadata:  make(metadata.MetadataMap),
			},
		}
	}
}
