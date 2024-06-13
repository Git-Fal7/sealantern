package types

type State int8

const (
	HANDSHAKING State = iota
	STATUS
	LOGIN
	PLAY
)
