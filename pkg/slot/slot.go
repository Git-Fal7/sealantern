package slot

import "github.com/seebs/nbt"

type SlotItem struct {
	ID         uint16
	Amount     uint8
	Durability uint16
	NBT        nbt.Compound
}

func (item SlotItem) DisplayName() nbt.String {
	displayNBT, ok := item.NBT["display"].(nbt.Compound)
	if !ok || displayNBT == nil {
		return ""
	}
	displayName, ok := displayNBT["Name"].(nbt.String)
	if !ok {
		return ""
	}
	return displayName
}

func (item SlotItem) Lore() []nbt.String {
	displayNBT, ok := item.NBT["display"].(nbt.Compound)
	if !ok || displayNBT == nil {
		return []nbt.String{}
	}
	loreList, ok := displayNBT["Lore"].(nbt.List)
	if !ok {
		return []nbt.String{}
	}
	lore, ok := loreList.GetStringList()
	if !ok {
		return []nbt.String{}
	}
	return lore
}

func (item *SlotItem) SetDisplayName(displayName string) {
	displayNBT, ok := item.NBT["display"].(nbt.Compound)
	if !ok || displayNBT == nil {
		displayNBT = nbt.Compound{}
	}
	displayNBT["Name"] = nbt.String(displayName)
	item.NBT["display"] = displayNBT
}

func (item *SlotItem) SetLore(lore ...string) {
	displayNBT, ok := item.NBT["display"].(nbt.Compound)
	if !ok || displayNBT == nil {
		displayNBT = nbt.Compound{}
	}
	lores := make([]nbt.String, len(lore))
	for i, s := range lore {
		lores[i] = nbt.String(s)
	}
	loreList := nbt.MakeStringList(lores)
	displayNBT["Lore"] = loreList
	item.NBT["display"] = displayNBT
}

func (item *SlotItem) SetCustomNBT(key string, tag nbt.Tag) {
	item.NBT[nbt.String(key)] = tag
}
