package scoreboard

import (
	"github.com/git-fal7/sealantern/minecraft/types"
)

type ScoreboardObjective struct {
	Name        string
	DisplayName string
	RenderType  types.ObjectiveRenderType
}
