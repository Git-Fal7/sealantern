package itemutil

import "github.com/git-fal7/sealantern/pkg/slot"

func IsEqual(i slot.SlotItem, i2 slot.SlotItem) bool {
	return i.Material.ID == i2.Material.ID && i.Durability == i2.Durability
}
