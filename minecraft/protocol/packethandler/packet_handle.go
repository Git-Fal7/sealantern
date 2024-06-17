package packethandler

import (
	"fmt"

	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/handler"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet/handlers"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/sealantern/server"
)

var (
	packets     map[int64]handler.SocketHandler = make(map[int64]handler.SocketHandler)
	playPackets map[int32]handler.PlayerHandler = make(map[int32]handler.PlayerHandler)
)

func packetTypeHash(state types.State, id int) int64 {
	return int64(id) ^ (int64(state) << 32)
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
	playPackets[0x0A] = &handlers.PlayAnimationServerHandler{
		Server: server,
	}
	playPackets[0x0E] = &handlers.PlayClickWindowHandler{
		Server: server,
	}
	playPackets[0x0B] = &handlers.PlayEntityActionHandler{
		Server: server,
	}
	playPackets[0x15] = &handlers.PlayClientSettingsHandler{}
	playPackets[0x17] = &handlers.PlayPluginMessageHandler{
		Server: server,
	}
}

func ExecutePacketHandler(conn *socket.Conn, packet protocol.Packet, id int, playerRegistry *player.PlayerRegistry) {
	if conn.State == types.PLAY {
		handler, ok := playPackets[int32(id)]
		if !ok {
			return
		}
		if playerRegistry == nil {
			return
		}
		if playerRegistry.GetPlayerFromUUID(conn.UUID) == nil {
			fmt.Println("Invalid player", conn.UUID)
			return
		}
		handler.Handle(playerRegistry.GetPlayerFromUUID(conn.UUID), packet)
	} else {
		packets[packetTypeHash(conn.State, int(id))].Handle(conn, packet)
	}
}
