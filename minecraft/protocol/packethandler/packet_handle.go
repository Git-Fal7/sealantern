package packethandler

import (
	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/handler"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet/handlers"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/sealantern/server"
)

var (
	packets     map[int16]handler.SocketHandler = make(map[int16]handler.SocketHandler)
	playPackets map[uint8]handler.PlayerHandler = make(map[uint8]handler.PlayerHandler)
)

func packetTypeHash(state types.State, id int) int16 {
	return int16(id) + (int16(state) << 8)
}

func InitRegistry(server server.Server) {
	packets[packetTypeHash(types.HANDSHAKING, 0x00)] = &handlers.HandshakeHandler{}
	packets[packetTypeHash(types.STATUS, 0x00)] = &handlers.StatusRequestHandler{
		Server: server,
	}
	packets[packetTypeHash(types.STATUS, 0x01)] = &handlers.StatusPingHandler{}
	packets[packetTypeHash(types.LOGIN, 0x00)] = &handlers.LoginStartHandler{
		Server: server,
	}

	playPackets[0x00] = &handlers.PlayKeepAliveHandler{}
	playPackets[0x01] = &handlers.PlayChatHandler{
		Server: server,
	}
	playPackets[0x02] = &handlers.PlayEntityUseHandler{
		Server: server,
	}
	playPackets[0x04] = &handlers.PlayPlayerPositionHandler{
		Server: server,
	}
	playPackets[0x05] = &handlers.PlayPlayerLookHandler{
		Server: server,
	}
	playPackets[0x06] = &handlers.PlayPlayerPositionAndLookHandler{
		Server: server,
	}
	playPackets[0x07] = &handlers.PlayPlayerDiggingHandler{
		Server: server,
	}
	playPackets[0x08] = &handlers.PlayBlockPlacementHandler{
		Server: server,
	}
	playPackets[0x09] = &handlers.PlayHeldItemChangeHandler{
		Server: server,
	}
	playPackets[0x0A] = &handlers.PlaySwingArmHandler{
		Server: server,
	}
	playPackets[0x0B] = &handlers.PlayEntityActionHandler{
		Server: server,
	}
	playPackets[0x0D] = &handlers.PlayCloseWindowHandler{}
	playPackets[0x0E] = &handlers.PlayClickWindowHandler{
		Server: server,
	}
	playPackets[0x14] = &handlers.PlayTabCompleteHandler{
		Server: server,
	}
	playPackets[0x15] = &handlers.PlayClientSettingsHandler{
		Server: server,
	}
	playPackets[0x17] = &handlers.PlayPluginMessageHandler{
		Server: server,
	}
}

func ExecutePacketHandler(conn *socket.Conn, packet protocol.Packet, id int, playerRegistry *player.PlayerRegistry) {
	if conn.State == types.PLAY {
		handler, ok := playPackets[uint8(id)]
		if !ok {
			return
		}
		if playerRegistry == nil {
			return
		}
		if playerRegistry.GetPlayerFromUUID(conn.UUID) == nil {
			return
		}
		handler.Handle(playerRegistry.GetPlayerFromUUID(conn.UUID), packet)
	} else {
		packets[packetTypeHash(conn.State, int(id))].Handle(conn, packet)
	}
}
