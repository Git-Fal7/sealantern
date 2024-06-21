package types

type GameStateReason uint8

const (
	GameStateReasonInvalidBed GameStateReason = iota
	GameStateReasonEndRaining
	GameStateReasonBeginRaining
	GameStateReasonChangeGamemode
	GameStateReasonEnterCredits
	GameStateReasonDemoMessage
	GameStateReasonArrowHittingPlayer
	GameStateReasonFadeValue
	GameStateReasonFadeTime
)
