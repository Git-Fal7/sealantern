package slot

import "github.com/seebs/nbt"

type SlotItem struct {
	ID         uint16
	Amount     uint8
	Durability uint16
	NBT        nbt.Compound
}
