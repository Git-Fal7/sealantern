package team

import (
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
}

func (team Team) GetPacket(mode types.TeamMode) *packet.PacketPlayTeams {
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
