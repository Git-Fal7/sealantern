package types

type Gamemode uint8

const (
	SURVIVAL Gamemode = iota
	CREATIVE
	ADVENTURE
	SPECTATOR
)
