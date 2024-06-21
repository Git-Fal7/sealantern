package npc

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/hologram"
	"github.com/google/uuid"
)

type NPCPlayer struct {
	entityID  uint16
	name      string
	uuid      uuid.UUID
	profile   profile.PlayerProfile
	position  world.Position
	holograms []*hologram.Hologram
}

func NewNPCPlayer(entityID uint16, position world.Position, properties []profile.Property) *NPCPlayer {
	b := make([]byte, 14)
	rand.Read(b)
	name := fmt.Sprintf("NPC-%x", b[2:6])
	npcUUID, _ := uuid.NewRandom()
	return &NPCPlayer{
		entityID: entityID,
		name:     name,
		uuid:     npcUUID,
		profile: profile.PlayerProfile{
			UUID:       npcUUID,
			Name:       name,
			Properties: properties,
		},
		position:  position,
		holograms: make([]*hologram.Hologram, 0),
	}
}

func (npc *NPCPlayer) SetText(offset float64, text ...string) {
	if len(npc.holograms) > len(text) {
		npc.holograms = npc.holograms[0:len(text)]
	} else if len(npc.holograms) < len(text) {
		for i := len(npc.holograms); i < len(text); i++ {
			npc.holograms = append(npc.holograms, hologram.NewHologram(npc.entityID+1000+uint16(i), "", world.Position{
				X: npc.position.X,
				Y: npc.position.Y + 1.5 + (offset * float64(len(text)-i)),
				Z: npc.position.Z,
			}))
		}
	}
	for i, hologram := range npc.holograms {
		hologram.DisplayName = text[i]
	}
}

func (npc NPCPlayer) EntityID() uint16 {
	return npc.entityID
}

func (npc NPCPlayer) Position() world.Position {
	return npc.position
}

func (npc NPCPlayer) SendPackets(p player.IPlayer) {
	p.WritePacket(&packet.PacketPlayPlayerListItem{
		Action: types.PlayerListActionAddPlayer,
		Entries: []types.PlayerListEntry{
			{
				Profile:     npc.profile,
				GameMode:    types.SURVIVAL,
				Ping:        0,
				DisplayName: nil,
			},
		},
	})
	p.WritePacket(&packet.PacketPlaySpawnPlayer{
		EntityID:       npc.entityID,
		PlayerUUID:     npc.uuid,
		PlayerPosition: npc.position,
		CurrentItem:    0, // Air
		Metadata: metadata.MetadataMap{
			metadata.MetadataPlayerSkinFlags: uint8(126),
		},
	})
	p.WritePacket(&packet.PacketPlayEntityHeadLook{
		EntityID: npc.entityID,
		HeadYaw:  npc.position.IntYaw(),
	})
	p.WritePacket(&packet.PacketPlayTeams{
		TeamName:          npc.name,
		Mode:              types.TeamModeCreate,
		FriendlyFire:      types.TeamFriendlyFireOff,
		NameTagVisibility: types.TeamNameTagVisibilityNever,
		Color:             0,
		Players: []string{
			npc.name,
		},
	})
	go func() {
		time.Sleep(time.Millisecond * 50)
		p.WritePacket(&packet.PacketPlayPlayerListItem{
			Action: types.PlayerListActionRemovePlayer,
			Entries: []types.PlayerListEntry{
				{
					Profile:     npc.profile,
					GameMode:    types.SURVIVAL,
					Ping:        0,
					DisplayName: nil,
				},
			},
		})
	}()
	for _, hologram := range npc.holograms {
		if hologram.DisplayName == "" {
			continue
		}
		for _, packet := range hologram.GetCreationPacket() {
			p.WritePacket(packet)
		}
	}
}

func (npc NPCPlayer) GetDestructionID() []uint16 {
	entitiesId := []uint16{
		npc.entityID,
	}
	for _, hologram := range npc.holograms {
		if hologram.DisplayName == "" {
			continue
		}
		entitiesId = append(entitiesId, hologram.EntityID())
	}
	return entitiesId
}
