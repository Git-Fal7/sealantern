package playerinventory

import (
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/pkg/slot"
)

type PlayerInventory struct {
	slots [45]slot.SlotItem
}

func NewPlayerInventory() *PlayerInventory {
	return &PlayerInventory{}
}

func (inv PlayerInventory) ID() uint8 {
	return 0
}

func (inv PlayerInventory) Title() string {
	return ""
}

func (inv PlayerInventory) Size() uint16 {
	return 45
}

func (inv *PlayerInventory) SetArmor(slot types.PlayerInventoryArmor, slotData slot.SlotItem) {
	if slot >= 4 {
		return
	}
	inv.slots[5+slot] = slotData // starts at index 5
}

func (inv *PlayerInventory) SetSlot(slot int, slotData slot.SlotItem) {
	if slot >= 36 {
		return
	}
	inv.slots[9+slot] = slotData // starts at index 9
}

func (inv *PlayerInventory) SetCrafting(slot int, slotData slot.SlotItem) {
	if slot >= 5 {
		return
	}
	inv.slots[slot] = slotData // starts at index 0
}

func (inv *PlayerInventory) SetHotbar(slot int, slotData slot.SlotItem) {
	if slot >= 9 {
		return
	}
	inv.slots[36+slot] = slotData // starts at index 36
}

func (inv PlayerInventory) Packets() []protocol.PacketOut {
	return []protocol.PacketOut{
		&packet.PacketPlayWindowItems{
			WindowID: 0,
			SlotData: inv.slots[:],
		},
	}
}
