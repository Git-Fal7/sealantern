package slot

import "github.com/seebs/nbt"

type SlotItem struct {
	ID         uint16
	Amount     uint8
	Durability uint16
	NBT        nbt.Compound
}

func (item SlotItem) DisplayName() string {
	displayNBT, ok := item.NBT["display"].(nbt.Compound)
	if !ok || displayNBT == nil {
		return ""
	}
	displayName, ok := displayNBT["Name"].(nbt.String)
	if !ok {
		return ""
	}
	return string(displayName)
}

func (item *SlotItem) SetDisplayName(displayName string) {
	displayNBT, ok := item.NBT["display"].(nbt.Compound)
	if !ok || displayNBT == nil {
		displayNBT = nbt.Compound{}
	}
	displayNBT["Name"] = nbt.String(displayName)
	item.NBT["display"] = displayNBT
}

func (item *SlotItem) SetCustomNBT(key string, tag nbt.Tag) {
	item.NBT[nbt.String(key)] = tag
}
