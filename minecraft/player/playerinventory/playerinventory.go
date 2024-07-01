package playerinventory

import (
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/pkg/slot"
)

type PlayerInventory struct {
	slots    [45]slot.SlotItem
	HeldSlot uint8
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
	return 36
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

func (inv *PlayerInventory) SetHeldItemSlot(slot uint8) {
	if slot >= 9 {
		return
	}
	inv.HeldSlot = slot
}

func (inv *PlayerInventory) GetHeldItem() slot.SlotItem {
	return inv.slots[36+inv.HeldSlot]
}

func (inv *PlayerInventory) GetArmorPackets(eid uint16) []protocol.PacketOut {
	return []protocol.PacketOut{
		&packet.PacketPlayEntityEquipment{
			EntityID: eid,
			Slot:     types.EquipimentSlotHelmet,
			Item:     inv.slots[5+types.PlayerInventoryArmorHelmet],
		},
		&packet.PacketPlayEntityEquipment{
			EntityID: eid,
			Slot:     types.EquipimentSlotChestplate,
			Item:     inv.slots[5+types.PlayerInventoryArmorChestplate],
		},
		&packet.PacketPlayEntityEquipment{
			EntityID: eid,
			Slot:     types.EquipimentSlotLeggings,
			Item:     inv.slots[5+types.PlayerInventoryArmorLeggings],
		},
		&packet.PacketPlayEntityEquipment{
			EntityID: eid,
			Slot:     types.EquipimentSlotBoots,
			Item:     inv.slots[5+types.PlayerInventoryArmorBoots],
		},
	}
}

func (inv *PlayerInventory) AddItem(item slot.SlotItem) {
	// Look for existing stacks
	toAdd := int(item.Amount)
	for i := 0; i < 36 || toAdd > 0; i++ {
		slotItem := inv.slots[9+i]
		if slotItem.ID != 0 {
			if slotItem.ID == item.ID && slotItem.Durability == item.Durability {
				space := 64 - int(slotItem.Amount)
				if space < 0 {
					continue
				}
				if space > toAdd {
					space = toAdd
				}
				slotItem.Amount += uint8(space)
				inv.slots[9+i] = slotItem
				toAdd -= space
			}
		}
	}

	// Look for empty slots
	if toAdd > 0 {
		for i := 0; i < 36 || toAdd > 0; i++ {
			slotItem := inv.slots[9+i]
			if slotItem.ID == 0 {
				var num int
				if toAdd > 64 {
					num = 64
				} else {
					num = toAdd
				}
				slotItem = item
				slotItem.Amount = uint8(num)
				inv.slots[9+i] = slotItem
				toAdd -= num
			}
		}
	}
}
