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

func (gui GUIInventory) ID() uint8 {
	return 1
}

func (gui GUIInventory) Title() string {
	return gui.title
}

func (gui GUIInventory) Size() uint16 {
	return uint16(len(gui.slots))
}

func (gui *GUIInventory) SetItem(slot slot.SlotItem, row int, column int) {
	if column+(row*9) >= len(gui.slots) {
		return
	}
	gui.slots[column+(row*9)] = slot
}

func (gui GUIInventory) Packets() []protocol.PacketOut {
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

func (gui *GUIInventory) Fill(slotItem slot.SlotItem) {
	for i := range gui.slots {
		gui.slots[i] = slotItem
	}
}

func (gui *GUIInventory) FillRow(row int, slotItem slot.SlotItem) {
	if row >= (len(gui.slots) / 9) {
		return
	}
	for i := 0; i < 9; i++ {
		gui.slots[i+(row*9)] = slotItem
	}
}

func (gui *GUIInventory) FillColumn(column int, slotItem slot.SlotItem) {
	if column >= 9 {
		return
	}
	for i := 0; i < (len(gui.slots) / 9); i++ {
		gui.slots[column+(i*9)] = slotItem
	}
}

func (gui *GUIInventory) FillBorder(slotItem slot.SlotItem) {
	gui.FillRow(0, slotItem)
	gui.FillRow((len(gui.slots)/9)-1, slotItem)
	gui.FillColumn(0, slotItem)
	gui.FillColumn(8, slotItem)
}
