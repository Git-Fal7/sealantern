package chunk

import (
	"github.com/git-fal7/sealantern/minecraft/blocks"
	"github.com/git-fal7/sealantern/minecraft/protocol/stream"
)

type BlockPalette interface {
	GetId(name string) int
	RecoverName(id int) string
	GetSize() int
	GetContent() []string
}

type ChunkBlockPalette struct {
	Map []string
}

type ChunkSection struct {
	Palette BlockPalette
	// This is a ridiculous way to put blocks in.. lmao
	Blocks     [4096]byte // 16 * 16 * 16
	SkyLight   [2048]byte
	BlockLight [2048]byte
}

type Chunk struct {
	ChunkX   int32
	ChunkZ   int32
	Sections [16]*ChunkSection // Max is 16
	Biomes   [256]byte
}

type ChunkKey struct {
	X int32
	Z int32
}

func (chunk *Chunk) GetSection(y int, skylight bool) *ChunkSection {
	if chunk.Sections[y] != nil {
		return chunk.Sections[y]
	}
	var blockLight [2048]byte
	for i := range blockLight {
		blockLight[i] = 0xFF
	}
	var skyLight [2048]byte
	if skylight {
		for i := range skyLight {
			skyLight[i] = 0xFF
		}
	}
	c := &ChunkSection{
		Palette: &ChunkBlockPalette{
			[]string{"minecraft:air"},
		},
		BlockLight: blockLight,
		SkyLight:   skyLight,
	}
	chunk.Sections[y] = c
	return c
}

func (chunk Chunk) ToData(skyLight bool) ([]byte, uint16) {
	return chunk.toData(skyLight, true)
}

func (chunk Chunk) toData(skyLight bool, entireChunk bool) ([]byte, uint16) {
	w := &stream.ProtocolWriter{}

	var bitmask uint16 = 0
	// Write blocks
	for i, s := range chunk.Sections {
		if s == nil ||
			(s.Palette.GetSize() == 1 && s.Palette.RecoverName(0) == "minecraft:air") {
			continue
		} else {
			bitmask |= 1 << i
		}

		for _, block := range s.Blocks {
			if block == 0 {
				w.WriteLittleEndianUInt16(0) // no need to convert and do all that for nothing.
			} else {
				value := uint16(blocks.BLOCK_REGISTRY.GetBlockId(s.Palette.RecoverName(int(block))))
				w.WriteLittleEndianUInt16(value)
			}
		}
	}

	for i, section := range chunk.Sections {
		if (bitmask & (1 << i)) == 0 {
			continue
		}
		w.WriteByteArray(section.BlockLight[:])
	}

	// Write sky lights
	if skyLight {
		for i, section := range chunk.Sections {
			if (bitmask & (1 << i)) == 0 {
				continue
			}
			w.WriteByteArray(section.SkyLight[:])
		}
	}

	// Write biomes
	if entireChunk {
		w.WriteByteArray(chunk.Biomes[:])
	}
	return w.Bytes(), bitmask // Last 3 bytes are useless and also breaks map chunk packet
}

func (palette *ChunkBlockPalette) GetId(name string) int {
	for i, v := range palette.Map {
		if v == name {
			return i
		}
	}
	palette.Map = append(palette.Map, name)
	return len(palette.Map) - 1
}

func (palette *ChunkBlockPalette) RecoverName(id int) string {
	if id < 0 || id >= len(palette.Map) {
		return "minecraft:air"
	}
	return palette.Map[id]
}

func (palette *ChunkBlockPalette) GetSize() int {
	return len(palette.Map)
}

func (palette *ChunkBlockPalette) GetContent() []string {
	return palette.Map
}

func (section *ChunkSection) SetBlock(x, y, z int, typ string) {
	section.Blocks[y<<8|z<<4|x] = byte(section.Palette.GetId(typ))
}
