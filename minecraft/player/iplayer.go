package player

import (
	"github.com/git-fal7/sealantern/minecraft/player/clientsettings"
	"github.com/git-fal7/sealantern/minecraft/player/playerinventory"
	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/chunk"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/gui"
	"github.com/git-fal7/sealantern/pkg/scoreboard/team"

	"github.com/google/uuid"
)

type IPlayer interface {
	UUID() uuid.UUID
	Username() string
	Profile() *profile.PlayerProfile
	ClientSettings() clientsettings.ClientSettings
	Position() world.Position
	SendMessage(msg component.IChatComponent)
	SendActionbar(msg component.IChatComponent)
	Disconnect(msg component.IChatComponent)
	WritePacket(packet protocol.PacketOut) error
	Active() bool
	ID() uint16
	KnownChunks() map[chunk.ChunkKey]bool
	Team() *team.Team
	SetTeam(team *team.Team)
	OpenInventory(inventory *gui.GUIInventory)
	PlayerInventory() *playerinventory.PlayerInventory
	UpdateInventory()
}
