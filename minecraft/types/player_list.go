package types

import (
	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/pkg/component"
)

type PlayerListEntry struct {
	Profile     profile.PlayerProfile
	GameMode    Gamemode
	Ping        int
	DisplayName component.IChatComponent
}

type PlayerListAction uint8

const (
	PlayerListActionAddPlayer PlayerListAction = iota
	PlayerListActionUpdateGamemode
	PlayerListActionUpdateLatency
	PlayerListActionUpdateDisplayName
	PlayerListActionRemovePlayer
)
