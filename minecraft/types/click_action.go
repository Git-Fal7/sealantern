package types

type ClickAction uint8

const (
	ClickActionLeftClickAir ClickAction = iota
	ClickActionLeftClickBlock
	ClickActionRightClickAir
	ClickActionRightClickBlock
)
