package clientsettings

import "github.com/git-fal7/sealantern/minecraft/types"

type ClientSettings struct {
	Locale             string
	ViewDistance       byte
	ChatMode           types.ChatMode
	ChatColors         bool
	DisplayedSkinParts types.DisplayedSkinParts
}
