package itemutil

import "github.com/git-fal7/sealantern/pkg/slot"

func IsEqual(i slot.SlotItem, i2 slot.SlotItem) bool {
	return i.ID == i2.ID && i.Durability == i2.Durability
}
