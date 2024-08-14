package handlers

import (
	"encoding/binary"
	"encoding/json"
	"strings"
	"time"

	"github.com/git-fal7/sealantern/config"
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
	"github.com/git-fal7/sealantern/pkg/inventory"
	"github.com/git-fal7/sealantern/pkg/itemutil"
	"github.com/git-fal7/sealantern/pkg/material"
	"github.com/git-fal7/sealantern/pkg/slot"
	"github.com/git-fal7/sealantern/pkg/uuidutil"
	"github.com/git-fal7/sealantern/sealantern/server"
	"github.com/google/uuid"
)

type HandshakeHandler struct{}

func (h *HandshakeHandler) Handle(p *socket.Conn, protoPacket protocol.Packet) {
	// set player state, address
	packet, _ := protoPacket.(*packet.PacketHandshake)
	if config.LanternConfig.InfoFowarding == config.InfoFowardingBungeeMode {
		split := strings.Split(packet.Address, "\x00")
		if (len(split) == 3 && split[2] != "") || len(split) == 4 {
			p.ProxyData = split
		}
	}
	p.State = packet.State
	p.Protocol = int32(packet.Protocol)
}

type StatusRequestHandler struct {
	Server server.Server
}

func (h *StatusRequestHandler) Handle(p *socket.Conn, protoPacket protocol.Packet) {
	serverListPing := &types.ServerListPing{
		Version: types.ServerListPingVersion{
			Name:     "SeaLantern",
			Protocol: 47,
		},
		Players: types.ServerListPingPlayers{
			MaxPlayers:    0,
			OnlinePlayers: 0,
			Sample:        make([]types.ServerListPingSample, 0),
		},
		Description: types.ServerListPingDescription{
			Motd: "",
		},
		Favicon: "",
		ModInfo: types.ServerListPingModInfo{
			ModType: "FML",
			ModList: make([]string, 0),
		},
	}
	serverListPingEvent := &events.ServerListPingEvent{
		ServerListPing: serverListPing,
		Allowed:        true,
	}
	h.Server.Event().Fire(serverListPingEvent)
	if !serverListPingEvent.Allowed {
		return
	}
	b, _ := json.Marshal(serverListPing)
	p.WritePacket(&packet.PacketStatusResponse{
		Response: string(b),
	})
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
	loginStartPacket, _ := protoPacket.(*packet.PacketLoginStart)
	if p.Protocol != 47 {
		println(p.Protocol)
		p.Disconnect(component.ChatMessage("Bad version no"))
		return
	}

	p.Username = loginStartPacket.Username
	var playerProfile *profile.PlayerProfile
	if config.LanternConfig.InfoFowarding == config.InfoFowardingBungeeMode {
		if p.ProxyData == nil {
			p.Disconnect(component.ChatMessage("If you wish to use IP forwarding, please enable it in your BungeeCord config as well!"))
			return
		}
		proxyData := p.ProxyData
		// 0: hostname
		// 1: spoofed addres
		p.UUID, _ = uuid.Parse(proxyData[2])
		playerProfile = &profile.PlayerProfile{
			UUID: p.UUID,
			Name: p.Username,
		}
		if len(proxyData) == 4 {
			// Spoofing properties
			var properties []profile.Property
			if err := json.Unmarshal([]byte(proxyData[3]), &properties); err != nil {
				p.Disconnect(component.ChatMessage("Unable to retrive properties"))
				return
			}
			playerProfile.Properties = properties
		}
	} else {
		p.UUID = uuidutil.GenerateOfflineUUID(p.Username)
		playerProfile = &profile.PlayerProfile{
			UUID: p.UUID,
			Name: p.Username,
		}
	}

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
		p.Disconnect(preLoginEvent.Reason)
	}

	player := connplayer.NewconnPlayer(playerProfile, p, h.Server.NextEID())
	h.Server.GetPlayerRegistry().RegisterPlayer(player)
	loginEvent := &events.PlayerPreJoinEvent{
		Player: player,
	}
	h.Server.Event().Fire(loginEvent)
	if loginEvent.Instance == nil {
		p.Disconnect(component.ChatMessage("No instances found for you"))
		return
	}
	selectedInstance := loginEvent.Instance

	permissionSetupEvent := &events.PermissionSetupEvent{
		Subject:     player,
		DefaultFunc: player.PermFunc,
	}
	h.Server.Event().Fire(permissionSetupEvent)
	player.PermFunc = permissionSetupEvent.Func()

	p.State = types.PLAY
	p.WritePacket(&packet.PacketPlayJoinGame{
		Gamemode:     selectedInstance.Gamemode,
		Dimension:    selectedInstance.World.Dimension,
		Difficulty:   selectedInstance.Difficulty,
		LevelType:    world.DEFAULT,
		MaxPlayers:   0xFF,
		ReducedDebug: false,
	})
	p.WritePacket(&packet.PacketPlayPlayerAbilities{
		Invulnerable: false,
		Fly:          false,
		CanFly:       false,
		Creative:     false,
		FlyingSpeed:  0.1,
		FieldOfView:  0.1,
	})
	player.SetGamemode(selectedInstance.Gamemode)
	h.Server.Event().Fire(&events.PlayerJoinEvent{
		Player: player,
	})
	err := selectedInstance.JoinPlayer(player)
	if err != nil {
		p.Disconnect(component.ChatMessage(err.Error()))
		return
	}
}

type PlayKeepAliveHandler struct{}

func (h *PlayKeepAliveHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	// set player's keepalive identifier
	packet, _ := protoPacket.(*packet.PacketPlayKeepAlive)
	if p.Conn.KeepAlive == packet.Identifier {
		p.Conn.KeepAlive = 0
	}
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
	if packet.Channel == "MC|Brand" {
		serverBrand := h.Server.Brand()
		buff := make([]byte, len(serverBrand)+1)
		length := binary.PutUvarint(buff, uint64(len(serverBrand)))
		copy(buff[length:], []byte(serverBrand))
	}

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
		p.SetSneaking(true)
		instance := h.Server.GetInstanceFromUUID(p.UUID())
		metaDataPacket := &packet.PacketPlayEntityMetadata{
			EntityID: p.ID(),
			Metadata: p.GetMetadata(),
		}
		for _, player := range instance.Players.GetPlayers() {
			if player.UUID() != p.UUID() {
				player.WritePacket(metaDataPacket)
			}
		}
	} else if actionPacket.ActionID == types.StopSneaking {
		p.SetSneaking(false)
		instance := h.Server.GetInstanceFromUUID(p.UUID())
		metaDataPacket := &packet.PacketPlayEntityMetadata{
			EntityID: p.ID(),
			Metadata: p.GetMetadata(),
		}
		for _, player := range instance.Players.GetPlayers() {
			if player.UUID() != p.UUID() {
				player.WritePacket(metaDataPacket)
			}
		}
	}
}

type PlaySwingArmHandler struct {
	Server server.Server
}

func (h *PlaySwingArmHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
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
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	// get victim id from instnace
	// the id exists in id, p is just for the attacker
	targetID := uint16(entityUsePacket.TargetID)
	victim := instance.Players.GetPlayerFromEID(targetID)
	if victim == nil {
		for _, npc := range instance.NPCs {
			if npc.EntityID() == targetID {
				h.Server.Event().Fire(&events.NpcInteractEvent{
					Player:       p,
					NPC:          npc,
					InteractType: entityUsePacket.Type,
				})
			}
		}
		return
	}
	if entityUsePacket.Type != types.UseEntityAttack {
		return
	}
	damageEvent := &events.PlayerDamageEvent{
		Attacker: p,
		Victim:   victim,
		Allowed:  true,
	}
	h.Server.Event().Fire(damageEvent)
}

type PlayClientSettingsHandler struct {
	Server server.Server
}

func (h *PlayClientSettingsHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	clientSettingsPacket, _ := protoPacket.(*packet.PacketPlayClientSettings)
	if p.Settings.DisplayedSkinParts != clientSettingsPacket.ClientSettings.DisplayedSkinParts {
		instance := h.Server.GetInstanceFromUUID(p.UUID())
		if instance != nil {
			p.WritePacket(&packet.PacketPlayEntityMetadata{
				EntityID: 0,
				Metadata: metadata.MetadataMap{
					metadata.MetadataPlayerSkinFlags: uint8(clientSettingsPacket.ClientSettings.DisplayedSkinParts),
				},
			})
			metadataPacket := &packet.PacketPlayEntityMetadata{
				EntityID: p.ID(),
				Metadata: metadata.MetadataMap{
					metadata.MetadataPlayerSkinFlags: uint8(clientSettingsPacket.ClientSettings.DisplayedSkinParts),
				},
			}
			for _, player := range instance.Players.GetPlayers() {
				if p.UUID() == player.UUID() {
					continue
				}
				player.WritePacket(metadataPacket)
			}
		}
	}
	p.Settings = clientSettingsPacket.ClientSettings
}

type PlayClickWindowHandler struct {
	Server server.Server
}

func (h *PlayClickWindowHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	clickWindowPacket, _ := protoPacket.(*packet.PacketPlayClickWindow)
	var clickedInventory inventory.Inventory
	if clickWindowPacket.WindowID == 0 {
		clickedInventory = p.Inventory
	} else {
		clickedInventory = p.OpenedInventory
	}
	if clickWindowPacket.Slot == 64537 {
		return // Dropping
	}
	clickEvent := &events.InventoryInteractEvent{
		Player:         p,
		Inventory:      clickedInventory,
		InteractedSlot: clickWindowPacket.Slot,
		InteractedItem: clickWindowPacket.ClickedItem,
		Allowed:        true,
	}
	h.Server.Event().Fire(clickEvent)
	p.WritePacket(&packet.PacketPlayConfirmTransaction{
		WindowID:     clickWindowPacket.WindowID,
		ActionNumber: clickWindowPacket.ActionNumber,
		Accepted:     clickEvent.Allowed,
	})
	if !clickEvent.Allowed {
		p.WritePacket(&packet.PacketPlaySetSlot{
			WindowID: 0xff,
			Slot:     -1,
			SlotData: p.ItemOnCursor,
		})
		p.UpdateInventory()
		return
	}
	if clickWindowPacket.WindowID == 0 {
		// Save stuff
		if clickWindowPacket.Mode == 0 {
			if p.Inventory.GetDirectSlot(int(clickWindowPacket.Slot)).Material != material.Air {
				p.Inventory.SetDirectSlot(int(clickWindowPacket.Slot), slot.SlotItem{Material: material.Air})
			}
			if p.ItemOnCursor.Material != material.Air {
				// process inventory
				p.Inventory.SetDirectSlot(int(clickWindowPacket.Slot), p.ItemOnCursor)
			}
			p.ItemOnCursor = clickWindowPacket.ClickedItem
		}
		if clickWindowPacket.Mode == 1 {
			if clickWindowPacket.ClickedItem.Material == material.Air {
				return
			}
			p.Inventory.SetDirectSlot(int(clickWindowPacket.Slot), slot.SlotItem{Material: material.Air})
			// Hotbar shift click
			toAdd := int(clickWindowPacket.ClickedItem.Amount)
			if clickWindowPacket.Slot >= 36 {
				for i := 9; i <= 35; i++ {
					item := p.Inventory.GetDirectSlot(i)
					if item.Material == material.Air {
						clickWindowPacket.ClickedItem.Amount = uint8(toAdd)
						p.Inventory.SetDirectSlot(i, clickWindowPacket.ClickedItem)
						toAdd -= toAdd
						break
					}
					if itemutil.IsEqual(item, clickWindowPacket.ClickedItem) {
						space := int(item.Material.GetStack()) - int(item.Amount)
						if space < 0 {
							continue
						}
						if space > toAdd {
							space = toAdd
						}
						item.Amount += uint8(space)
						p.Inventory.SetDirectSlot(i, item)
						toAdd -= space
						if toAdd <= 0 {
							break
						}
					}
				}
				if toAdd > 0 {
					clickWindowPacket.ClickedItem.Amount = uint8(toAdd)
					p.Inventory.SetDirectSlot(int(clickWindowPacket.Slot), clickWindowPacket.ClickedItem)
				}
			} else {
				for i := 36; i <= 44; i++ {

					item := p.Inventory.GetDirectSlot(i)
					if item.Material == material.Air {
						clickWindowPacket.ClickedItem.Amount = uint8(toAdd)
						p.Inventory.SetDirectSlot(i, clickWindowPacket.ClickedItem)
						toAdd -= toAdd
						break
					}
					if itemutil.IsEqual(item, clickWindowPacket.ClickedItem) {
						space := int(item.Material.GetStack()) - int(item.Amount)
						if space < 0 {
							continue
						}
						if space > toAdd {
							space = toAdd
						}
						item.Amount += uint8(space)
						p.Inventory.SetDirectSlot(i, item)
						toAdd -= space
						if toAdd <= 0 {
							break
						}
					}
				}
				if toAdd > 0 {
					clickWindowPacket.ClickedItem.Amount = uint8(toAdd)
					p.Inventory.SetDirectSlot(int(clickWindowPacket.Slot), clickWindowPacket.ClickedItem)
				}
			}
		}
		if clickWindowPacket.Mode == 2 {
			selectedItem := p.Inventory.GetDirectSlot(int(clickWindowPacket.Slot))
			movedItem := p.Inventory.GetDirectSlot(36 + int(clickWindowPacket.Button))
			p.Inventory.SetDirectSlot(int(clickWindowPacket.Slot), movedItem)
			p.Inventory.SetDirectSlot(36+int(clickWindowPacket.Button), selectedItem)
		}
	}
}

type PlayCloseWindowHandler struct{}

func (h *PlayCloseWindowHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	closeWindowPacket, _ := protoPacket.(*packet.PacketPlayCloseWindow)
	if closeWindowPacket.WindowID != 0 {
		p.OpenedInventory = nil
	}
	p.ItemOnCursor = slot.SlotItem{Material: material.Air}
}

type PlayPlayerDiggingHandler struct {
	Server server.Server
}

func (h *PlayPlayerDiggingHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	diggingPacket, _ := protoPacket.(*packet.PacketPlayPlayerDigging)
	if diggingPacket.Status == types.FinishedDigging {
		// Break event
		instance := h.Server.GetInstanceFromUUID(p.UUID())
		block := instance.World.GetBlock(diggingPacket.Location.X, diggingPacket.Location.Y, diggingPacket.Location.Z)
		breakEvent := &events.PlayerBreakBlockEvent{
			Player:   p,
			Block:    block,
			Location: diggingPacket.Location,
			Allowed:  true,
		}
		h.Server.Event().Fire(breakEvent)
		if !breakEvent.Allowed {
			p.WritePacket(&packet.PacketPlayBlockChange{
				Location: diggingPacket.Location,
				Type:     block,
			})
			return
		}
		blockChangePacket := &packet.PacketPlayBlockChange{
			Location: diggingPacket.Location,
			Type:     0,
		}
		instance.World.SetBlock(diggingPacket.Location.X, diggingPacket.Location.Y, diggingPacket.Location.Z, "minecraft:air", true)
		for _, player := range instance.Players.GetPlayers() {
			player.WritePacket(blockChangePacket)
		}
	} else if diggingPacket.Status == types.ShootArrowOrFinishEating {
		instance := h.Server.GetInstanceFromUUID(p.UUID())
		heldItem := p.Inventory.GetHeldItem()
		if material.IsSword(heldItem.Material) || heldItem.Material == material.Bow {
			// Send blocking packet
			p.SetBlocking(false)
			metadataPacket := &packet.PacketPlayEntityMetadata{
				EntityID: p.ID(),
				Metadata: p.GetMetadata(),
			}
			if heldItem.Material == material.Bow {
				if p.PlayerInventory().HasItem(material.Arrow) {
					shootEvent := &events.PlayerShootBowEvent{
						Player:  p,
						Force:   p.BowCharge,
						Allowed: true,
					}
					h.Server.Event().Fire(shootEvent)
					if !shootEvent.Allowed {
						p.UpdateInventory()
					}
				}
				p.BowCharge = 0
			}
			for _, player := range instance.Players.GetPlayers() {
				if player.UUID() != p.UUID() {
					player.WritePacket(metadataPacket)
				}
			}
		}
	}
}

type PlayBlockPlacementHandler struct {
	Server server.Server
}

func (h *PlayBlockPlacementHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	blockPlacementPacket, _ := protoPacket.(*packet.PacketPlayBlockPlacement)
	action := types.ClickActionRightClickBlock
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	var blockIDAt int
	if blockPlacementPacket.Face == 0xff {
		previous := p.LastPlacementPacket
		if previous == nil || previous.HeldItem.Material != blockPlacementPacket.HeldItem.Material {
			action = types.ClickActionRightClickAir
			blockIDAt = -1
		} else {
			p.LastPlacementPacket = nil
			return
		}
	} else {
		blockIDAt = instance.World.GetBlock(blockPlacementPacket.Location.X, blockPlacementPacket.Location.Y, blockPlacementPacket.Location.Z)
	}
	if material.IsSword(blockPlacementPacket.HeldItem.Material) || (blockPlacementPacket.HeldItem.Material == material.Bow && p.PlayerInventory().HasItem(material.Arrow)) {
		// Send blocking packet
		// Deep copying the metadata
		removeBlockingMetadata := make(metadata.MetadataMap)
		for k, v := range p.GetMetadata() {
			if k == metadata.MetadataIndexStatus {
				var status uint8 = 0
				status, _ = v.(uint8)
				removeBlockingMetadata[k] = status & ^metadata.StatusFlagBlocking
			}
			removeBlockingMetadata[k] = v
		}
		removeBlockingPacket := &packet.PacketPlayEntityMetadata{
			EntityID: p.ID(),
			Metadata: removeBlockingMetadata,
		}

		p.SetBlocking(true)
		metadataPacket := &packet.PacketPlayEntityMetadata{
			EntityID: p.ID(),
			Metadata: p.GetMetadata(),
		}
		p.BowCharge = 0
		go func() {
			for p.BowCharge < 1 && p.IsBlocking() {
				p.BowCharge += 0.10
				time.Sleep(time.Millisecond * 100)
			}
		}()
		for _, player := range instance.Players.GetPlayers() {
			if player.UUID() != p.UUID() {
				player.WritePacket(removeBlockingPacket)
				go func(player *connplayer.ConnectedPlayer) {
					// might redo.
					time.Sleep(50 * time.Millisecond)
					player.WritePacket(metadataPacket)
				}(player)
			}
		}
	}
	p.LastPlacementPacket = blockPlacementPacket
	timer := time.NewTimer(time.Millisecond * 100) // 2 ticks
	go func() {
		<-timer.C
		if p.LastPlacementPacket == blockPlacementPacket {
			p.LastPlacementPacket = nil
		}
	}()
	// TODO: Placaement
	h.Server.Event().Fire(&events.PlayerInteractEvent{
		Player:          p,
		Slot:            p.PlayerInventory().GetHeldItem(),
		Action:          action,
		BlockAt:         blockIDAt,
		BlockAtLocation: blockPlacementPacket.Location,
	})
}

type PlayHeldItemChangeHandler struct {
	Server server.Server
}

func (h *PlayHeldItemChangeHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	heldItemChange, _ := protoPacket.(*packet.PacketPlayHeldItemChange)
	slotChangeEvent := &events.PlayerHeldItemChangeEvent{
		Player:       p,
		PreviousSlot: p.PlayerInventory().HeldSlot,
		CurrentSlot:  heldItemChange.Slot,
		Allowed:      true,
	}
	p.SetBlocking(false)
	p.BowCharge = 0
	h.Server.Event().Fire(slotChangeEvent)
	if !slotChangeEvent.Allowed {
		p.WritePacket(&packet.PacketPlayHeldItemChange{
			Slot: p.PlayerInventory().HeldSlot,
		})
		return
	}
	p.PlayerInventory().SetHeldItemSlot(heldItemChange.Slot)
	instance := h.Server.GetInstanceFromUUID(p.UUID())
	entityEquipmentPacket := &packet.PacketPlayEntityEquipment{
		EntityID: p.ID(),
		Slot:     types.EquipmentSlotHeldItem,
		Item:     p.PlayerInventory().GetHeldItem(),
	}
	for _, player := range instance.Players.GetPlayers() {
		if player.UUID() != p.UUID() {
			player.WritePacket(entityEquipmentPacket)
		}
	}
}

type PlayTabCompleteHandler struct {
	Server server.Server
}

func (h *PlayTabCompleteHandler) Handle(p *connplayer.ConnectedPlayer, protoPacket protocol.Packet) {
	tabCompletePacket, _ := protoPacket.(*packet.PacketPlayTabCompleteServer)
	if len(tabCompletePacket.Text) == 0 {
		return
	}
	// TODO: fire tab complete event
	splitString := strings.Split(tabCompletePacket.Text, "")
	if splitString[0] != "/" {
		return
	}
	splitString = splitString[1:] // remove the first slash
	splitCmd := strings.Split(strings.Join(splitString, ""), " ")
	cmdText := splitCmd[0]
	args := splitCmd[1:]
	if len(args) == 0 {
		return
	}
	cmd, err := h.Server.Command().GetCommand(cmdText)
	if err != nil {
		// not found
		return
	}
	suggestions := cmd.Suggest(command.SimpleCommandInvocation{
		Arguments: args,
		Source:    p,
	})
	p.WritePacket(&packet.PacketPlayTabComplete{
		Matches: suggestions,
	})
}
