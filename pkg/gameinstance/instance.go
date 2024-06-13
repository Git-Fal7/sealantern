package gameinstance

import (
	"fmt"

	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/player/connplayer"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/map_world"

	"github.com/google/uuid"
)

type GameInstance struct {
	World      *map_world.Map
	Gamemode   types.Gamemode
	Difficulty world.Difficulty
	Players    *player.PlayerRegistry
}

func (instance *GameInstance) JoinPlayer(p *connplayer.ConnectedPlayer) error {
	if instance.HasPlayerFromUUID(p.UUID()) {
		return fmt.Errorf("player already joined this instance")
	}
	p.WritePacket(&packet.PacketPlayServerDifficulty{
		Difficulty: instance.Difficulty,
	})
	p.Teleport(instance.World.Spawn)
	instance.World.SendChunksAroundPlayer(p)
	instance.Players.RegisterPlayer(p)
	entries := make([]types.PlayerListEntry, 0)
	players := instance.Players.GetPlayers()
	for _, player := range players {
		entries = append(entries, types.PlayerListEntry{
			Profile:  *player.Profile(),
			GameMode: types.SURVIVAL,
			Ping:     0,
		})
	}
	packetPlayerListItem := &packet.PacketPlayPlayerListItem{
		Action:  types.PlayerListActionAddPlayer,
		Entries: entries,
	}
	packetSpawnJoinedPlayer := &packet.PacketPlaySpawnPlayer{
		EntityID:       p.ID(),
		PlayerUUID:     p.UUID(),
		PlayerPosition: p.Position(),
		CurrentItem:    0,
	}
	packetHeadRotation := &packet.PacketPlayEntityHeadLook{
		EntityID: p.ID(),
		HeadYaw:  p.Position().IntYaw(),
	}
	p.WritePacket(packetPlayerListItem)
	if p.CurrentTeam != nil {
		p.WritePacket(p.CurrentTeam.GetPacket(types.TeamModeCreate))
	}
	for _, player := range players {
		if player.UUID() != p.UUID() {
			player.WritePacket(packetPlayerListItem)
			player.WritePacket(packetSpawnJoinedPlayer)
			player.WritePacket(packetHeadRotation)
			p.WritePacket(&packet.PacketPlaySpawnPlayer{
				EntityID:       player.ID(),
				PlayerUUID:     player.UUID(),
				PlayerPosition: player.Position(),
				CurrentItem:    0,
			})
			p.WritePacket(&packet.PacketPlayEntityHeadLook{
				EntityID: player.ID(),
				HeadYaw:  player.Position().IntYaw(),
			})
			if player.CurrentTeam != nil {
				p.WritePacket(player.CurrentTeam.GetPacket(types.TeamModeCreate))
			}
			if p.CurrentTeam != nil {
				player.WritePacket(p.CurrentTeam.GetPacket(types.TeamModeRemove))
				player.WritePacket(p.CurrentTeam.GetPacket(types.TeamModeCreate))
			}
		}
	}

	/*
		p.WritePacket(&packet.PacketPlayTeams{
			TeamName:          "01",
			Mode:              types.TeamModeCreate,
			TeamDisplayName:   "",
			TeamPrefix:        "&c",
			TeamSuffix:        "",
			FriendlyFire:      types.TeamFriendlyFireOff,
			NameTagVisibility: types.TeamNameTagVisibilityAlways,
			Color:             0,
			Players: []string{
				p.Username(),
			},
		})
		for _, player := range players {
			if player.UUID() == p.UUID() {
				continue
			}
			p.WritePacket(&packet.PacketPlayTeams{
				TeamName:          "02",
				Mode:              types.TeamModeCreate,
				TeamDisplayName:   "",
				TeamPrefix:        "&d",
				TeamSuffix:        "",
				FriendlyFire:      types.TeamFriendlyFireOff,
				NameTagVisibility: types.TeamNameTagVisibilityAlways,
				Color:             0,
				Players: []string{
					player.Username(),
				},
			})
			player.WritePacket(&packet.PacketPlayTeams{
				TeamName: "01",
				Mode:     types.TeamModeAddPlayer,
				Players: []string{
					p.Username(),
				},
			})
		}
	*/
	return nil
}

func (instance *GameInstance) QuitPlayer(p *connplayer.ConnectedPlayer) error {
	if !instance.HasPlayerFromUUID(p.UUID()) {
		return fmt.Errorf("player doesnt exist in this instance")
	}
	instance.Players.UnregisterPlayer(p)
	entries := make([]types.PlayerListEntry, 0)
	entries = append(entries, types.PlayerListEntry{
		Profile: *p.Profile(),
	})
	packetPlayerListItem := &packet.PacketPlayPlayerListItem{
		Action:  types.PlayerListActionRemovePlayer,
		Entries: entries,
	}
	packetDestroyEntities := &packet.PacketPlayDestroyEntites{
		EntityIDs: make([]uint16, 0),
	}
	packetDestroyEntities.EntityIDs = append(packetDestroyEntities.EntityIDs, p.ID())
	for _, player := range instance.Players.GetPlayers() {
		player.WritePacket(packetPlayerListItem)
		player.WritePacket(packetDestroyEntities)
	}
	return nil
}

func (instance *GameInstance) MovePlayer(p *connplayer.ConnectedPlayer, from world.Position, to world.Position, onGround bool) {
	var movePacket protocol.PacketOut
	var headRotationPacket protocol.PacketOut
	if from.X == to.X && from.Y == to.Y && from.Z == to.Z {
		movePacket = &packet.PacketPlayEntityLook{
			EntityID: p.ID(),
			Yaw:      to.IntYaw(),
			Pitch:    to.IntPitch(),
			OnGround: onGround,
		}
		headRotationPacket = &packet.PacketPlayEntityHeadLook{
			EntityID: p.ID(),
			HeadYaw:  to.IntYaw(),
		}
	} else if from.Yaw == to.Yaw && from.Pitch == to.Pitch {
		movePacket = &packet.PacketPlayEntityRelativeMove{
			EntityID: p.ID(),
			DeltaX:   int8(to.IntX() - from.IntX()),
			DeltaY:   int8(to.IntY() - from.IntY()),
			DeltaZ:   int8(to.IntZ() - from.IntZ()),
			OnGround: onGround,
		}
	} else {
		movePacket = &packet.PacketPlayEntityLookAndRelativeMove{
			EntityID: p.ID(),
			DeltaX:   int8(to.IntX() - from.IntX()),
			DeltaY:   int8(to.IntY() - from.IntY()),
			DeltaZ:   int8(to.IntZ() - from.IntZ()),
			Yaw:      to.IntYaw(),
			Pitch:    to.IntPitch(),
			OnGround: onGround,
		}
		headRotationPacket = &packet.PacketPlayEntityHeadLook{
			EntityID: p.ID(),
			HeadYaw:  to.IntYaw(),
		}
	}
	for _, player := range instance.Players.GetPlayers() {
		if player.UUID() != p.UUID() {
			player.WritePacket(movePacket)
			if headRotationPacket != nil {
				player.WritePacket(headRotationPacket)
			}
		}
	}
}

func (instance *GameInstance) HasPlayerFromUUID(uuid uuid.UUID) bool {
	return instance.Players.GetPlayerFromUUID(uuid) != nil
}

// basically load chunks around player
func (instance *GameInstance) Tick() {
	for _, player := range instance.Players.GetPlayers() {
		instance.World.SendChunksAroundPlayer(player)
	}
}
