package types

type Animation uint8

const (
	SwingArm Animation = iota
	TakeDamage
	LeaveBed
	EatFood
	CriticalEffect
	MagicCriticalEffect
)
