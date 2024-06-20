package team

import (
	"slices"
	"sync"

	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
)

type Team struct {
	Name              string
	DisplayName       string
	Prefix            string
	Suffix            string
	FriendlyFire      types.TeamFriendlyFire
	NameTagVisibility types.TeamNameTagVisibility
	Color             int8
	Players           []string
	mutex             sync.RWMutex
}

func (team *Team) AddPlayer(name string) {
	team.mutex.Lock()
	team.Players = append(team.Players, name)
	defer team.mutex.Unlock()
}

func (team *Team) RemovePlayer(name string) {
	team.mutex.Lock()
	team.Players = slices.DeleteFunc(team.Players, func(username string) bool {
		return username == name
	})
	defer team.mutex.Unlock()
}

func (team *Team) GetPacket(mode types.TeamMode) *packet.PacketPlayTeams {
	return &packet.PacketPlayTeams{
		TeamName:          team.Name,
		Mode:              mode,
		TeamDisplayName:   team.DisplayName,
		TeamPrefix:        team.Prefix,
		TeamSuffix:        team.Suffix,
		FriendlyFire:      team.FriendlyFire,
		NameTagVisibility: team.NameTagVisibility,
		Color:             team.Color,
		Players:           team.Players,
	}
}
