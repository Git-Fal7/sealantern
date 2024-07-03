package types

type TitleAction uint8

const (
	TitleActionSetTitle TitleAction = iota
	TitleActionSetSubtitle
	TitleActionSetTimesAndDisplay
	TitleActionHide
	TitleActionReset
)
