package types

type ScoreboardMode int8

const (
	CreateScoreboard ScoreboardMode = iota
	RemoveScoreboard
	UpdateScoreboard
)

type UpdateScoreAction int8

const (
	CreateScoreItem UpdateScoreAction = iota
	RemoveScoreItem
)

type ObjectiveDisplaySlot int8

const (
	ScoreboardPositionPlayerList ObjectiveDisplaySlot = iota
	ScoreboardPositionSidebar
	ScoreboardPositionBelowName
)

type ObjectiveRenderType string

const (
	ObjectiveRenderTypeInteger ObjectiveRenderType = "integer"
	ObjectiveRenderTypeHearts  ObjectiveRenderType = "hearts"
)
