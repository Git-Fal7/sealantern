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
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/inventory"
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
	eid                 uint16
	Sneaking            bool
	Health              float32
	Invincibile         bool
	Settings            clientsettings.ClientSettings
	CurrentTeam         *team.Team
	KnownChunkKeys      map[chunk.ChunkKey]bool
	OpenedInventory     inventory.Inventory
	Inventory           *playerinventory.PlayerInventory
	LastPlacementPacket *packet.PacketPlayBlockPlacement
	ItemOnCursor        slot.SlotItem
}

func NewconnPlayer(profile *profile.PlayerProfile, conn *socket.Conn, eid uint16) *ConnectedPlayer {
	return &ConnectedPlayer{
		profile:         profile,
		Conn:            conn,
		eid:             eid,
		Health:          20,
		Invincibile:     false,
		KnownChunkKeys:  make(map[chunk.ChunkKey]bool),
		OpenedInventory: nil,
		Inventory:       playerinventory.NewPlayerInventory(),
		ItemOnCursor:    slot.SlotItem{ID: 0},
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

func (p *ConnectedPlayer) ID() uint16 {
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
	if p.Sneaking {
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
