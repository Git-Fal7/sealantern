package playerinventory

import (
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/pkg/slot"
)

type PlayerInventory struct {
	slots    [36]slot.SlotItem
	armor    [4]slot.SlotItem
	crafting [5]slot.SlotItem
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
	inv.armor[slot] = slotData
}

func (inv *PlayerInventory) SetSlot(slot int, slotData slot.SlotItem) {
	if slot >= 36 {
		return
	}
	inv.slots[slot] = slotData
}

func (inv *PlayerInventory) SetCrafting(slot int, slotData slot.SlotItem) {
	if slot >= 5 {
		return
	}
	inv.crafting[slot] = slotData
}

func (inv PlayerInventory) Packets() []protocol.PacketOut {
	slots := []slot.SlotItem{}
	slots = append(slots, inv.crafting[:]...)
	slots = append(slots, inv.armor[:]...)
	slots = append(slots, inv.slots[:]...)
	return []protocol.PacketOut{
		&packet.PacketPlayWindowItems{
			WindowID: 0,
			SlotData: slots,
		},
	}
}
