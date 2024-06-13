package command

import (
	"github.com/git-fal7/sealantern/pkg/component"
	"github.com/git-fal7/sealantern/pkg/permission"
)

type CommandSource interface {
	permission.Subject
	SendMessage(msg component.IChatComponent)
}
