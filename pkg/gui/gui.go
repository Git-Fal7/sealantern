package gui

import (
	"github.com/git-fal7/sealantern/minecraft/protocol"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/types"
	"github.com/git-fal7/sealantern/pkg/slot"
)

type GUIInventory struct {
	title string
	slots []slot.SlotItem
}

func NewGUIInventory(title string, rows int) *GUIInventory {
	return &GUIInventory{
		title: title,
		slots: make([]slot.SlotItem, rows*9),
	}
}

func (gui *GUIInventory) SetItem(slot slot.SlotItem, slotIndex uint8) {
	if slotIndex >= uint8(len(gui.slots)) {
		return
	}
	gui.slots[slotIndex] = slot
}

func (gui GUIInventory) GetCreationPacket() []protocol.PacketOut {
	packets := []protocol.PacketOut{
		&packet.PacketPlayOpenWindow{
			WindowID:      1,
			WindowType:    types.WindowTypeChest,
			WindowTitle:   gui.title,
			NumberOfSlots: uint8(len(gui.slots)),
		},
		&packet.PacketPlayWindowItems{
			WindowID: 1,
			SlotData: gui.slots,
		},
	}
	return packets
}