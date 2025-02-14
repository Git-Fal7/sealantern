package player

import (
	"github.com/git-fal7/sealantern/minecraft/entity"
	"github.com/git-fal7/sealantern/minecraft/player/clientsettings"
	"github.com/git-fal7/sealantern/minecraft/player/playerinventory"
	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/chunk"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/inventory"
	"github.com/git-fal7/sealantern/pkg/scoreboard/team"
)

type IPlayer interface {
	entity.LivingEntity
	Username() string
	Profile() *profile.PlayerProfile
	ClientSettings() clientsettings.ClientSettings
	Position() world.Position
	SendMessage(msg component.IChatComponent)
	SendActionbar(msg component.IChatComponent)
	Disconnect(msg component.IChatComponent)
	WritePacket(packet protocol.PacketOut) error
	KnownChunks() map[chunk.ChunkKey]bool
	Team() *team.Team
	SetTeam(team *team.Team)
	OpenInventory(inventory inventory.Inventory)
	PlayerInventory() *playerinventory.PlayerInventory
	UpdateInventory()
	SetTablistHeader(header component.IChatComponent)
	SetTablistFooter(footer component.IChatComponent)
	SetTablistHeaderFooter(header component.IChatComponent, footer component.IChatComponent)
	SendTitle(title component.IChatComponent, subtitle component.IChatComponent, fadein int32, stay int32, fadeout int32)
	PlaySound(Location world.BlockPosition, Sound types.SoundEffect, volume float32, pitch float32)
	FoodLevel() int
	SetFoodLevel(foodLevel int)
	Saturation() float32
	SetSaturation(saturation float32)
	Gamemode() types.Gamemode
	SetGamemode(gamemode types.Gamemode)
	GetMetadata() metadata.MetadataMap
	SetSneaking(sneaking bool)
	IsSneaking() bool
	SetBlocking(blocking bool)
	IsBlocking() bool
}
