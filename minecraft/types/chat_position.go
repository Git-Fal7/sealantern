package types

type ChatPosition uint8

const (
	CHAT_BOX ChatPosition = iota
	SYSTEM
	ACTION_BAR
)
