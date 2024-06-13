package types

type UseEntityType uint8

const (
	UseEntityInteract UseEntityType = iota
	UseEntityAttack
	UseEntityInteractAt
)
