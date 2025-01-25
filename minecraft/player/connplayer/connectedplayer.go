package connplayer

import (
	"github.com/git-fal7/sealantern/minecraft/player/clientsettings"
	"github.com/git-fal7/sealantern/minecraft/player/playerinventory"
	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/chunk"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/inventory"
	"github.com/git-fal7/sealantern/pkg/material"
	"github.com/git-fal7/sealantern/pkg/permission"
	"github.com/git-fal7/sealantern/pkg/scoreboard/team"
	"github.com/git-fal7/sealantern/pkg/slot"

	"github.com/google/uuid"
)

type ConnectedPlayer struct {
	Conn                *socket.Conn
	profile             *profile.PlayerProfile
	Pos                 world.Position
	PermFunc            permission.Func
	eid                 int32
	health              float32
	foodLevel           int
	saturation          float32
	Settings            clientsettings.ClientSettings
	CurrentTeam         *team.Team
	KnownChunkKeys      map[chunk.ChunkKey]bool
	OpenedInventory     inventory.Inventory
	Inventory           *playerinventory.PlayerInventory
	LastPlacementPacket *packet.PacketPlayBlockPlacement
	ItemOnCursor        slot.SlotItem
	gamemode            types.Gamemode
	metadata            metadata.MetadataMap
	BowCharge           float32
}

func NewconnPlayer(profile *profile.PlayerProfile, conn *socket.Conn, eid int32) *ConnectedPlayer {
	return &ConnectedPlayer{
		profile:         profile,
		Conn:            conn,
		eid:             eid,
		health:          20,
		foodLevel:       20,
		saturation:      5,
		KnownChunkKeys:  make(map[chunk.ChunkKey]bool),
		OpenedInventory: nil,
		Inventory:       playerinventory.NewPlayerInventory(),
		ItemOnCursor:    slot.SlotItem{Material: material.Air},
		gamemode:        types.SURVIVAL,
		metadata:        make(metadata.MetadataMap),
	}
}

func (p *ConnectedPlayer) UUID() uuid.UUID {
	return p.profile.UUID
}

func (p *ConnectedPlayer) Username() string {
	return p.profile.Name
}
func (p *ConnectedPlayer) Profile() *profile.PlayerProfile {
	return p.profile
}

func (p *ConnectedPlayer) SendMessage(message component.IChatComponent) {
	p.WritePacket(&packet.PacketPlayMessage{
		Component: message,
		Position:  types.CHAT_BOX,
	})
}

func (p *ConnectedPlayer) SendActionbar(message component.IChatComponent) {
	p.WritePacket(&packet.PacketPlayMessage{
		Component: message,
		Position:  types.ACTION_BAR,
	})
}

func (p *ConnectedPlayer) Disconnect(message component.IChatComponent) {
	p.Conn.Disconnect(message)
}

func (p *ConnectedPlayer) WritePacket(packet protocol.PacketOut) error {
	return p.Conn.WritePacket(packet)
}

func (p *ConnectedPlayer) Teleport(position world.Position) {
	p.WritePacket(&packet.PacketPlayPlayerPositionAndLookClient{
		Position: position,
		Flags:    0x00,
	})
	p.Pos = position
}
func (p *ConnectedPlayer) Position() world.Position {
	return p.Pos
}

func (p *ConnectedPlayer) ID() int32 {
	return p.eid
}

/* permission.Subject */
func (p *ConnectedPlayer) HasPermission(permission string) bool {
	return p.PermissionValue(permission).Bool()
}

func (p *ConnectedPlayer) PermissionValue(permission string) permission.TriState {
	return p.PermFunc(permission)
}

/* Else */
func (p *ConnectedPlayer) GetEyePosition() world.Position {
	var pos world.Position
	if p.IsSneaking() {
		//		return 1.54
		pos = p.Position().Add(world.Position{
			X: 0,
			Y: 1.54,
			Z: 0,
		})
	} else {
		//		return 1.62
		pos = p.Position().Add(world.Position{
			X: 0,
			Y: 1.62,
			Z: 0,
		})
	}
	return pos
}

func (p *ConnectedPlayer) KnownChunks() map[chunk.ChunkKey]bool {
	return p.KnownChunkKeys
}

func (p *ConnectedPlayer) ClientSettings() clientsettings.ClientSettings {
	return p.Settings
}

func (p *ConnectedPlayer) Team() *team.Team {
	return p.CurrentTeam
}

func (p *ConnectedPlayer) SetTeam(team *team.Team) {
	team.AddPlayer(p.Username())
	p.CurrentTeam = team
}

func (p *ConnectedPlayer) PlayerInventory() *playerinventory.PlayerInventory {
	return p.Inventory
}

func (p *ConnectedPlayer) OpenInventory(inventory inventory.Inventory) {
	p.OpenedInventory = inventory
	for _, packet := range inventory.Packets() {
		p.WritePacket(packet)
	}
}

func (p *ConnectedPlayer) UpdateInventory() {
	if p.OpenedInventory != nil {
		p.WritePacket(p.OpenedInventory.Packets()[1])
	}
	p.WritePacket(p.Inventory.Packets()[0])
}

func (p *ConnectedPlayer) SetTablistHeader(header component.IChatComponent) {
	p.WritePacket(&packet.PacketPlayerListHeaderFooter{
		Header: header,
	})
}

func (p *ConnectedPlayer) SetTablistFooter(footer component.IChatComponent) {
	p.WritePacket(&packet.PacketPlayerListHeaderFooter{
		Footer: footer,
	})
}

func (p *ConnectedPlayer) SetTablistHeaderFooter(header component.IChatComponent, footer component.IChatComponent) {
	p.WritePacket(&packet.PacketPlayerListHeaderFooter{
		Header: header,
		Footer: footer,
	})
}

func (p *ConnectedPlayer) SendTitle(title component.IChatComponent, subtitle component.IChatComponent, fadein int32, stay int32, fadeout int32) {
	if title == nil && subtitle == nil {
		return
	}
	if title != nil {
		p.WritePacket(&packet.PacketPlayTitle{
			Action: types.TitleActionSetTitle,
			Title:  title,
		})
	}
	if subtitle != nil {
		p.WritePacket(&packet.PacketPlayTitle{
			Action:   types.TitleActionSetSubtitle,
			Subtitle: subtitle,
		})
	}
	p.WritePacket(&packet.PacketPlayTitle{
		Action:  types.TitleActionSetTimesAndDisplay,
		FadeIn:  fadein,
		Stay:    stay,
		FadeOut: fadeout,
	})
}

func (p *ConnectedPlayer) PlaySound(location world.BlockPosition, sound types.SoundEffect, volume float32, pitch float32) {
	p.WritePacket(&packet.PacketPlaySoundEffect{
		SoundName:      sound,
		EffectPosition: location,
		Volume:         volume,
		Pitch:          pitch,
	})
}

func (p *ConnectedPlayer) sendHealth() {
	p.WritePacket(&packet.PacketPlayUpdateHealth{
		Health:         p.health,
		Food:           p.foodLevel,
		FoodSaturation: p.saturation,
	})
}

func (p *ConnectedPlayer) Health() float32 {
	return p.health
}

func (p *ConnectedPlayer) SetHealth(health float32) {
	p.health = health
	p.sendHealth()
}

func (p *ConnectedPlayer) FoodLevel() int {
	return p.foodLevel
}

func (p *ConnectedPlayer) SetFoodLevel(foodLevel int) {
	p.foodLevel = min(foodLevel, 20)
	p.sendHealth()
}

func (p *ConnectedPlayer) Saturation() float32 {
	return p.saturation
}

func (p *ConnectedPlayer) SetSaturation(saturation float32) {
	p.saturation = min(saturation, float32(p.foodLevel))
	p.sendHealth()
}

func (p *ConnectedPlayer) Gamemode() types.Gamemode {
	return p.gamemode
}

func (p *ConnectedPlayer) SetGamemode(gamemode types.Gamemode) {
	if p.gamemode != gamemode {
		p.gamemode = gamemode
		p.WritePacket(&packet.PacketPlayChangeGameState{
			Reason: types.GameStateReasonChangeGamemode,
			Value:  float32(p.gamemode),
		})
	}
}

func (p *ConnectedPlayer) GetMetadata() metadata.MetadataMap {
	return p.metadata
}

func (p *ConnectedPlayer) SetSneaking(sneaking bool) {
	var status uint8 = 0
	status, _ = p.metadata[metadata.MetadataIndexStatus].(uint8)
	if sneaking {
		p.metadata[metadata.MetadataIndexStatus] = status | metadata.StatusFlagSneaking
	} else {
		p.metadata[metadata.MetadataIndexStatus] = status & ^metadata.StatusFlagSneaking
	}
}

func (p *ConnectedPlayer) IsSneaking() bool {
	var status uint8 = 0
	status, _ = p.metadata[metadata.MetadataIndexStatus].(uint8)
	return status&metadata.StatusFlagSneaking != 0
}

func (p *ConnectedPlayer) SetBlocking(blocking bool) {
	var status uint8 = 0
	status, _ = p.metadata[metadata.MetadataIndexStatus].(uint8)
	if blocking {
		p.metadata[metadata.MetadataIndexStatus] = status | metadata.StatusFlagBlocking
	} else {
		p.metadata[metadata.MetadataIndexStatus] = status & ^metadata.StatusFlagBlocking
	}
}

func (p *ConnectedPlayer) IsBlocking() bool {
	var status uint8 = 0
	status, _ = p.metadata[metadata.MetadataIndexStatus].(uint8)
	return status&metadata.StatusFlagBlocking != 0
}
