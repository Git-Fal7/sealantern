package events

import (
	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/gameinstance"
	"github.com/git-fal7/sealantern/pkg/npc"
	"github.com/git-fal7/sealantern/pkg/permission"
)

type PlayerPreLoginResult uint8

const (
	AllowedPreLogin PlayerPreLoginResult = iota
	DeniedPreLogin
)

type PlayerPreLoginEvent struct {
	Connection     *socket.Conn
	PreLoginResult PlayerPreLoginResult
	Reason         component.StringDisconnectComponent
}

func (e *PlayerPreLoginEvent) Deny(component component.StringDisconnectComponent) {
	e.PreLoginResult = DeniedPreLogin
	e.Reason = component
}

func (e *PlayerPreLoginEvent) Allow(component component.StringDisconnectComponent) {
	e.PreLoginResult = AllowedPreLogin
	e.Reason = component
}

type PlayerPreJoinEvent struct {
	Player   player.IPlayer
	Instance *gameinstance.GameInstance
}

func (e *PlayerPreJoinEvent) SetInstance(instance *gameinstance.GameInstance) {
	e.Instance = instance
}

type PlayerJoinEvent struct {
	Player player.IPlayer
}

type PlayerQuitEvent struct {
	Player player.IPlayer
}

type PlayerKickEvent struct {
	Player player.IPlayer
	Reason string
}

type PlayerChatEvent struct {
	Player  player.IPlayer
	Message string
}

type PlayerClickType byte

const (
	PlayerRightClick PlayerClickType = iota
	PlayerLeftClick
)

type PlayerInteractEvent struct {
	Player    player.IPlayer
	ClickType PlayerClickType
}

type PlayerMoveEvent struct {
	Player       player.IPlayer
	FromPosition world.Position
	ToPosition   world.Position
	OnGround     bool
	Allowed      bool
}

type PluginMessageEvent struct {
	Channel string
	Data    []byte
}

type PermissionSetupEvent struct {
	Subject     permission.Subject
	DefaultFunc permission.Func
	fn          permission.Func
}

func (p *PermissionSetupEvent) Func() permission.Func {
	if p.fn == nil {
		return p.DefaultFunc
	}
	return p.fn
}

func (p *PermissionSetupEvent) SetFunc(fn permission.Func) {
	if fn == nil {
		return
	}
	p.fn = fn
}

type PlayerDamageEvent struct {
	Victim   player.IPlayer
	Attacker player.IPlayer
	Allowed  bool
}

func (e *PlayerDamageEvent) SetAllowed(allowed bool) {
	e.Allowed = allowed
}

type ServerListPingEvent struct {
	ServerListPing *types.ServerListPing
	Allowed        bool
}

func (e *ServerListPingEvent) SetAllowed(allowed bool) {
	e.Allowed = allowed
}

type NpcInteractEvent struct {
	Player       player.IPlayer
	NPC          npc.NPC
	InteractType types.UseEntityType
}
