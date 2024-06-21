package scoreboard

import (
	"fmt"
	"strings"

	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/pkg/scoreboard/team"
)

type ScoreboardSidebar struct {
	ScoreboardObjective
	Lines   []*ScoreboardSidebarLine
	Viewers []player.IPlayer
}

func NewSidebar(name string) *ScoreboardSidebar {
	return &ScoreboardSidebar{
		ScoreboardObjective: ScoreboardObjective{
			Name:        name,
			DisplayName: "",
			RenderType:  types.ObjectiveRenderTypeInteger,
		},
		Viewers: make([]player.IPlayer, 0),
	}
}
func (sidebar *ScoreboardSidebar) UpdateLinesString(lines ...string) {
	sidebar.UpdateLines(lines)
}

func (sidebar *ScoreboardSidebar) UpdateLines(lines []string) {
	if len(lines) > 15 || len(lines) == 0 {
		return
	}
	if len(sidebar.Lines) > len(lines) {
		// Delete old extra lines
		for i := len(sidebar.Lines) - 1; i >= len(lines); i-- {
			sidebar.sendToViewers(sidebar.Lines[i].GetScorePacket(i, types.RemoveScoreItem, sidebar.Name))
			sidebar.sendToViewers(sidebar.Lines[i].GetPacket(types.TeamModeRemove))
		}
		sidebar.Lines = sidebar.Lines[0:len(lines)]
	} else if len(sidebar.Lines) < len(lines) {
		// Add extra lines
		for i := len(sidebar.Lines); i < len(lines); i++ {
			sidebar.Lines = append(sidebar.Lines, &ScoreboardSidebarLine{
				Team: team.Team{
					Name:              fmt.Sprintf("__fakeScore%d", i),
					DisplayName:       "",
					FriendlyFire:      types.TeamFriendlyFireOff,
					NameTagVisibility: types.TeamNameTagVisibilityAlways,
					Color:             0,
				},
			})
		}
	}
	for i := 0; i < len(lines); i++ {
		sidebar.UpdateLine(i, lines[i])
	}
}

func (sidebar *ScoreboardSidebar) UpdateLine(index int, line string) {
	if index > 14 || index < 0 || len(line) > 48 {
		return
	}
	team := sidebar.Lines[index]
	mergedText := fmt.Sprintf("%s%s%s", team.Prefix, team.Content, team.Suffix)
	if mergedText == line { // Ignore line if it doesn't change
		return
	}
	sidebar.sendToViewers(team.GetScorePacket(index, types.RemoveScoreItem, sidebar.Name))
	sidebar.sendToViewers(team.GetPacket(types.TeamModeRemove))
	splitLine := strings.Split(line, "")
	if len(splitLine) <= 16 {
		team.Prefix = ""
		team.Content = strings.Join(splitLine[:], "")
		team.Players = []string{team.Content}
		team.Suffix = ""
	} else if len(splitLine) <= 32 {
		team.Prefix = strings.Join(splitLine[0:16], "")
		team.Content = strings.Join(splitLine[16:], "")
		team.Players = []string{team.Content}
		team.Suffix = ""
	} else { // <= 48
		team.Prefix = strings.Join(splitLine[0:16], "")
		team.Content = strings.Join(splitLine[16:32], "")
		team.Players = []string{team.Content}
		team.Suffix = strings.Join(splitLine[32:], "")
	}
	sidebar.sendToViewers(team.GetPacket(types.TeamModeCreate))
	sidebar.sendToViewers(team.GetScorePacket(index, types.CreateScoreItem, sidebar.Name))
}

func (sidebar *ScoreboardSidebar) AddViewer(player player.IPlayer) {
	sidebar.Viewers = append(sidebar.Viewers, player)

	// Create objective to player
	player.WritePacket(&packet.PacketPlayScoreboardObjective{
		ObjectiveName: sidebar.Name,
		DisplayName:   sidebar.DisplayName,
		Mode:          types.CreateScoreboard,
		RenderType:    types.ObjectiveRenderTypeInteger,
	})
	player.WritePacket(&packet.PacketPlayDisplayScoreboard{
		Position:  types.ScoreboardPositionSidebar,
		ScoreName: sidebar.Name,
	})
	for i, line := range sidebar.Lines {
		player.WritePacket(line.GetPacket(types.TeamModeCreate))
		player.WritePacket(line.GetScorePacket(i, types.CreateScoreItem, sidebar.Name))
	}
}

func (sidebar *ScoreboardSidebar) SetTitle(title string) {
	if len(title) > 32 {
		return
	}
	sidebar.DisplayName = title
	sidebar.sendToViewers(&packet.PacketPlayScoreboardObjective{
		ObjectiveName: sidebar.Name,
		DisplayName:   sidebar.DisplayName,
		Mode:          types.UpdateScoreboard,
		RenderType:    types.ObjectiveRenderTypeInteger,
	})
}

func (sidebar ScoreboardSidebar) sendToViewers(packet protocol.PacketOut) {
	for _, viewer := range sidebar.Viewers {
		viewer.WritePacket(packet)
	}
}

type ScoreboardSidebarLine struct {
	team.Team
	Content string
}

func (team *ScoreboardSidebarLine) GetScorePacket(line int, action types.UpdateScoreAction, objective string) *packet.PacketPlayUpdateScore {
	return &packet.PacketPlayUpdateScore{
		ScoreName:     team.Content,
		Action:        action,
		ObjectiveName: objective,
		Value:         15 - line,
	}
}
