package handler

import (
	"github.com/git-fal7/sealantern/minecraft/player/connplayer"
	"github.com/git-fal7/sealantern/minecraft/player/socket"
	"github.com/git-fal7/sealantern/minecraft/protocol"
)

type SocketHandler interface {
	Handle(p *socket.Conn, packet protocol.Packet)
}

type PlayerHandler interface {
	Handle(p *connplayer.ConnectedPlayer, packet protocol.Packet)
}
