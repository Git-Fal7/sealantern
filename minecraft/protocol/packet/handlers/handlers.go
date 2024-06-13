package handlers

import (
	"crypto/md5"
	"strings"
	"time"

	"github.com/git-fal7/sealantern/minecraft/player/connplayer"
	"github.com/git-fal7/sealantern/minecraft/player/profile"
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/metadata"
	"github.com/git-fal7/sealantern/pkg/command"
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/events"
	"github.com/git-fal7/sealantern/pkg/rayutil"
	"github.com/git-fal7/sealantern/sealantern/server"
)

type HandshakeHandler struct{}

func (h *HandshakeHandler) Handle(p *socket.Conn, protoPacket protocol.Packet) {
	// set player state, address
	packet, _ := protoPacket.(*packet.PacketHandshake)
	p.State = packet.State
	p.Protocol = int32(packet.Protocol)
}

type StatusRequestHandler struct {
	Server server.Server
}

func (h *StatusRequestHandler) Handle(p *socket.Conn, protoPacket protocol.Packet) {
	/*
		func (packet *PacketStatusRequest) Handle(player *Player) {
			protocol := COMPATIBLE_PROTO[0]
			if IsCompatible(player.protocol) {
				protocol = player.protocol
			}

			max_players := config.TyphoonConfig.MaxPlayers
			motd := config.TyphoonConfig.Motd

			count := player.core.playerRegistry.GetPlayerCount()
			if max_players < count && !config.TyphoonConfig.Restricted {
				max_players = count
			}

			response := PacketStatusResponse{
				Response: fmt.Sprintf(`{"version":{"name":"Typhoon","protocol":%d},"players":{"max":%d,"online":%d,"sample":[]},"description":{"text":"%s"},"favicon":"%s","modinfo":{"type":"FML","modList":[]}}`, protocol, max_players, count, JsonEscape(motd), JsonEscape(config.TyphoonConfig.Favicon)),
			}
			player.WritePacket(&response)
		}
	*/
	// todo: motd, maxplayers
	response := &packet.PacketStatusResponse{
		Response: `{"version":{"name":"Typhoon","protocol":47},"players":{"max":0,"online":0,"sample":[]},"description":{"text":"Big smonk"},"favicon":"","modinfo":{"type":"FML","modList":[]}}`,
	}
	p.WritePacket(response)
}

type StatusPingHandler struct {
}

func (h *StatusPingHandler) Handle(p *socket.Conn, protoPacket protocol.Packet) {
	packet, _ := protoPacket.(*packet.PacketStatusPing)
	p.WritePacket(packet)
}

type LoginStartHandler struct {
	Server server.Server
}

func (h *LoginStartHandler) Handle(p *socket.Conn, protoPacket protocol.Packet) {
	lPacket, _ := protoPacket.(*packet.PacketLoginStart)
	if p.Protocol != 47 {
		p.Disconnect(&component.StringDisconnectComponent{
			Text: "Bad version no",
		})
		return
	}
	// kick if "full"

	p.Username = lPacket.Username
	// Offline mode
	md5uuid := md5.Sum([]byte("OfflinePlayer:" + p.Username))
	md5uuid[6] = (md5uuid[6] & 0x0f) | uint8((3&0xf)<<4)
	md5uuid[8] = (md5uuid[8] & 0x3f) | 0x80
	p.UUID = md5uuid
	// check if compression is on
	p.WritePacket(&packet.PacketLoginSuccess{
		UUID:     p.UUID,
		Username: p.Username,
	})
	preLoginEvent := &events.PlayerPreLoginEvent{
		Connection:     p,
		PreLoginResult: events.AllowedPreLogin,
	}
	h.Server.Event().Fire(preLoginEvent)
	if preLoginEvent.PreLoginResult == events.DeniedPreLogin {
		p.Disconnect(&preLoginEvent.Reason)
	}

	profile := &profile.PlayerProfile{
		UUID: p.UUID,
		Name: p.Username,
	}
	player := connplayer.NewconnPlayer(profile, p, h.Server.NextEID())
	h.Server.GetPlayerRegistry().RegisterPlayer(player)
	loginEvent := &events.PlayerPreJoinEvent{
		Player: player,
	}
	h.Server.Event().Fire(loginEvent)
	if loginEvent.Instance == nil {
		p.Disconnect(&component.StringDisconnectComponent{
			Text: "No instances found for you",
		})
		return
	}

	permissionSetupEvent := &events.PermissionSetupEvent{
		Subject:     player,
		DefaultFunc: player.PermFunc,
	}
	h.Server.Event().Fire(permissionSetupEvent)
	player.PermFunc = permissionSetupEvent.Func()

	p.State = types.PLAY
	p.WritePacket(&packet.PacketPlayJoinGame{
		//EntityId:     0,
		Gamemode:     loginEvent.Instance.Gamemode,
		Dimension:    loginEvent.Instance.World.Dimension,
		Difficulty:   loginEvent.Instance.Difficulty,
		LevelType:    world.DEFAULT,
		MaxPlayers:   0xFF,
		ReducedDebug: false,
	})
	p.WritePacket(&packet.PacketPlayPlayerAbilities{
		Invulnerable: false,
		Fly:          false,
		CanFly:       true,
		Creative:     false,
		FlyingSpeed:  0.1,
		FieldOfView:  0.1,
	})
	err := loginEvent.Instance.JoinPlayer(player)
	if err != nil {
		p.Disconnect(&component.StringChatComponent{
			Text: err.Error(),
		})
		return
	}
	h.Server.Event().Fire(&events.PlayerJoinEvent{
		Player: player,
	})
}

type PlayKeepAliveHandler struct{}

func (h *PlayKeepAliveHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	// set player's keepalive identifier
	packet, _ := protoPacket.(*packet.PacketPlayKeepAlive)
	p.Conn.KeepAlive = packet.Identifier
}

type PlayChatHandler struct {
	Server server.Server
}

func (h *PlayChatHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	packet, _ := protoPacket.(*packet.PacketPlayChat)
	if len(packet.Message) <= 0 {
		return
	}
	if packet.Message[0] != '/' {
		// Call chat event
		h.Server.Event().Fire(&events.PlayerChatEvent{
			Player:  p,
			Message: packet.Message,
		})
	} else {
		// Treat message as command.
		if len(packet.Message) == 1 || strings.HasPrefix(packet.Message, "/ ") {
			return
		}
		parsedCmd := strings.Split(packet.Message, " ")
		cmdName := parsedCmd[0][1:]
		var args []string = nil
		if len(parsedCmd) > 1 {
			args = parsedCmd[1:]
		}
		cmd, err := h.Server.Command().GetCommand(cmdName)
		if err != nil {
			return
		}
		cmd.Execute(command.SimpleCommandInvocation{
			Arguments: args,
			Source:    p,
		})
	}
}

type PlayPluginMessageHandler struct {
	Server server.Server
}

func (h *PlayPluginMessageHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	packet, _ := protoPacket.(*packet.PacketPlayPluginMessage)
	/*
		if packet.Channel == "MC|Brand" || packet.Channel == "minecraft:brand" {
			log.Printf("%s is using %s client", player.name, string(packet.Data))
			buff := make([]byte, len(player.core.brand)+1)
			length := binary.PutUvarint(buff, uint64(len(player.core.brand)))
			copy(buff[length:], []byte(player.core.brand))
		}
	*/
	h.Server.Event().Fire(&events.PluginMessageEvent{
		Channel: packet.Channel,
		Data:    packet.Data,
	})
}

type PlayPlayerPositionAndLookHandler struct {
	Server server.Server
}

func (h *PlayPlayerPositionAndLookHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	packet, _ := protoPacket.(*packet.PacketPlayPlayerPositionAndLook)
	moveEvent := &events.PlayerMoveEvent{
		Player:       p,
		FromPosition: p.Pos,
		ToPosition:   packet.Position,
		OnGround:     packet.OnGround,
		Allowed:      true,
	}
	h.Server.Event().Fire(moveEvent)
	if !moveEvent.Allowed {
		// teleport back
		return
	}
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	if instance == nil {
		return
	}
	p.Pos = packet.Position
	instance.MovePlayer(p, moveEvent.FromPosition, moveEvent.ToPosition, packet.OnGround)
}

type PlayPlayerPositionHandler struct {
	Server server.Server
}

func (h *PlayPlayerPositionHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	packet, _ := protoPacket.(*packet.PacketPlayPlayerPosition)
	position := world.Position{
		X:     packet.X,
		Y:     packet.FeetY,
		Z:     packet.Z,
		Yaw:   p.Position().Yaw,
		Pitch: p.Position().Pitch,
	}
	moveEvent := &events.PlayerMoveEvent{
		Player:       p,
		FromPosition: p.Pos,
		ToPosition:   position,
		OnGround:     packet.OnGround,
		Allowed:      true,
	}
	h.Server.Event().Fire(moveEvent)
	if !moveEvent.Allowed {
		// teleport back
		return
	}
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	if instance == nil {
		return
	}
	p.Pos = position
	instance.MovePlayer(p, moveEvent.FromPosition, moveEvent.ToPosition, packet.OnGround)
}

type PlayPlayerLookHandler struct {
	Server server.Server
}

func (h *PlayPlayerLookHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	packet, _ := protoPacket.(*packet.PacketPlayPlayerLook)
	position := world.Position{
		X:     p.Position().X,
		Y:     p.Position().Y,
		Z:     p.Position().Z,
		Yaw:   packet.Yaw,
		Pitch: packet.Pitch,
	}
	moveEvent := &events.PlayerMoveEvent{
		Player:       p,
		FromPosition: p.Pos,
		ToPosition:   position,
		OnGround:     packet.OnGround,
		Allowed:      true,
	}
	h.Server.Event().Fire(moveEvent)
	if !moveEvent.Allowed {
		// teleport back
		return
	}
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	if instance == nil {
		return
	}
	p.Pos = position
	instance.MovePlayer(p, moveEvent.FromPosition, moveEvent.ToPosition, packet.OnGround)
}

type PlayEntityActionHandler struct {
	Server server.Server
}

func (h *PlayEntityActionHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	actionPacket, _ := protoPacket.(*packet.PacketPlayEntityAction)
	if actionPacket.ActionID == types.StartSneaking {
		//p.Sneaking = true
		instance := h.Server.GetInstanceFromUUID(p.UUID())
		metaDataPacket := &packet.PacketPlayEntityMetadata{
			EntityID: p.ID(),
			Metadata: metadata.MetadataMap{
				metadata.MetadataIndexStatus: metadata.StatusFlagSneaking,
			},
		}
		for _, player := range instance.Players.GetPlayers() {
			if player.UUID() != p.UUID() {
				player.WritePacket(metaDataPacket)
			}
		}
	} else if actionPacket.ActionID == types.StopSneaking {
		instance := h.Server.GetInstanceFromUUID(p.UUID())
		metaDataPacket := &packet.PacketPlayEntityMetadata{
			EntityID: p.ID(),
			Metadata: metadata.MetadataMap{
				metadata.MetadataIndexStatus: metadata.MetadataIndexStatus.Index & ^metadata.StatusFlagSneaking,
			},
		}
		for _, player := range instance.Players.GetPlayers() {
			if player.UUID() != p.UUID() {
				player.WritePacket(metaDataPacket)
			}
		}
	}
}

type PlayAnimationServerHandler struct {
	Server server.Server
}

func (h *PlayAnimationServerHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	animationPacket := &packet.PacketPlayAnimationClient{
		EntityID:  p.ID(),
		Animation: types.SwingArm,
	}
	for _, player := range instance.Players.GetPlayers() {
		if player.UUID() != p.UUID() {
			player.WritePacket(animationPacket)
		}
	}
}

type PlayEntityUseHandler struct {
	Server server.Server
}

func (h *PlayEntityUseHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	entityUsePacket, _ := protoPacket.(*packet.PacketPlayUseEntity)
	if entityUsePacket.Type != types.UseEntityAttack {
		return
	}
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	// get victim id from instnace
	// the id exists in id, p is just for the attacker
	targetID := uint16(entityUsePacket.TargetID)
	victim := instance.Players.GetPlayerFromEID(targetID)
	if victim == nil {
		return
	}
	if victim.Invincibile {
		return
	}
	damageEvent := &events.PlayerDamageEvent{
		Attacker: p,
		Victim:   victim,
		Allowed:  true,
	}
	h.Server.Event().Fire(damageEvent)
	if !damageEvent.Allowed {
		return
	}
	animationPacket := &packet.PacketPlayAnimationClient{
		EntityID:  targetID,
		Animation: types.TakeDamage,
	}
	distance := rayutil.GetRayBetween(victim.Position(), p.Position())
	rayLength := rayutil.GetVelocityRay(distance).Normalize()
	rayLength = rayLength.Multiply(world.Vector{
		X: ((0.5 + 1) / 3.0),
		Y: 0,
		Z: ((0.5 + 1) / 3.0),
	})
	rayLength.Y = 0.45
	for _, player := range instance.Players.GetPlayers() {
		player.WritePacket(animationPacket)
	}
	victim.WritePacket(&packet.PacketPlayEntityVelocity{
		EntityID: 0,
		Velocity: rayLength,
	})
	victim.Invincibile = true
	timer := time.NewTimer(time.Millisecond * 400)
	go func() {
		<-timer.C
		victim.Invincibile = false
	}()

	victim.Health = victim.Health - 0.5
	victim.WritePacket(&packet.PacketPlayUpdateHealth{
		Health:         victim.Health,
		Food:           20,
		FoodSaturation: 5,
	})
}

type PlayClientSettingsHandler struct {
}

func (h *PlayClientSettingsHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	packet, _ := protoPacket.(*packet.PacketPlayClientSettings)
	p.Settings = packet.ClientSettings
}
