package server

import (
	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/pkg/command"
	"github.com/git-fal7/sealantern/pkg/gameinstance"

	"github.com/google/uuid"
	"github.com/robinbraemer/event"
)

type Server interface {
	GetPlayerRegistry() *player.PlayerRegistry
	Event() event.Manager
	Command() *command.Manager
	GetInstanceFromUUID(uuid.UUID) *gameinstance.GameInstance
	Brand() string
}
