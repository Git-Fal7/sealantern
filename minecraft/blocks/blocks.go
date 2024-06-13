package blocks

var (
	BLOCK_REGISTRY = generateRegistry()
)

type BlockRegistry struct {
	counter    uint16
	NameToGuid map[string]uint16
	GuidToName map[uint16]string
}

func generateRegistry() *BlockRegistry {
	r := &BlockRegistry{
		counter:    0,
		NameToGuid: make(map[string]uint16),
		GuidToName: make(map[uint16]string),
	}

	for _, v := range GetLegacyMapping() {
		r.GetGuid(v.Name)
	}

	return r
}

func (registry *BlockRegistry) GetGuid(name string) uint16 {
	if val, ok := registry.NameToGuid[name]; ok {
		return val
	} else {
		nid := registry.counter
		registry.counter++
		registry.NameToGuid[name] = nid
		registry.GuidToName[nid] = name
		return nid
	}
}

func (registry *BlockRegistry) GetName(guid uint16) string {
	if val, ok := registry.GuidToName[guid]; ok {
		return val
	} else {
		return "minecraft:stone"
	}
}

func (registry *BlockRegistry) GetBlockId(name string) int {
	return registry.GetLegacyBlockId(name)
}

func (registry *BlockRegistry) GetLegacyBlockId(name string) int {
	block := GetLegacyFromName(name)
	for block.Protocol != 0 && block.Protocol > 47 && block.Fallback != nil {
		block = GetLegacyFromName(*block.Fallback)
	}
	return block.GetBlockState()
}

func (registry *BlockRegistry) GetLegacyBlockTypeData(name string) (int, int) {
	block := GetLegacyFromName(name)
	for block.Protocol != 0 && block.Protocol > 47 && block.Fallback != nil {
		block = GetLegacyFromName(*block.Fallback)
	}
	return block.Id, block.Data
}
