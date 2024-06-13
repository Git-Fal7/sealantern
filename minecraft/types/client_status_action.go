package types

type ClientStatusAction uint8

const (
	PERFORM_RESPAWN ClientStatusAction = iota
	REQUEST_STATS
	OPEN_INVENTORY
)
