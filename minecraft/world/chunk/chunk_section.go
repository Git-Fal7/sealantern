package chunk

type ChunkSection struct {
	Palette    BlockPalette
	Blocks     [4096]byte // 16 * 16 * 16
	SkyLight   [2048]byte
	BlockLight [2048]byte
}

func (section *ChunkSection) SetBlock(x, y, z int, id uint16) {
	sID := section.Palette.GetId(id)
	section.Blocks[y<<8|z<<4|x] = byte(sID)
}
