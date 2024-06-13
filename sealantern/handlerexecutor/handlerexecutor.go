package handlerexecutor

import (
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/protocol"
)

type HandlerExecutor interface {
	ExecuteHandler(conn *socket.Conn, packet protocol.Packet)
}
