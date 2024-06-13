package types

type EntityAction uint8

const (
	StartSneaking EntityAction = iota
	StopSneaking
	EnttiyActionLeaveBed
	StartSprinting
	StopSprinting
	JumpWithHorse
	OpenHorseInventory
)
