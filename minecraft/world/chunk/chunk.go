package chunk

import (
	"github.com/git-fal7/sealantern/minecraft/protocol/stream"
)

type Chunk struct {
	ChunkX   int32
	ChunkZ   int32
	data []byte
	bitmask uint16
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
			[]uint16{0},
		},
		BlockLight: blockLight,
		SkyLight:   skyLight,
	}
	chunk.Sections[y] = c
	return c
}

func (chunk *Chunk) Reload(skyLight bool) {
	data, bitmask := chunk.getData(skyLight, true)
	chunk.data = data
	chunk.bitmask = bitmask
}

func (chunk Chunk) Data() ([]byte, uint16) {
	return chunk.data, chunk.bitmask
}

func (chunk Chunk) getData(skyLight bool, entireChunk bool) ([]byte, uint16) {
	w := &stream.ProtocolWriter{}

	var bitmask uint16 = 0

	// Write blocks
	for i, s := range chunk.Sections {
		if s == nil {
			continue
		}
		sectionPalette := s.Palette.GetContent()
		if len(sectionPalette) == 1 && sectionPalette[0] == 0 {
			continue
		}
		bitmask |= 1 << i
		for _, sectionID := range s.Blocks {
			if sectionID == 0 {
				w.WriteLittleEndianUInt16(0) // no need to convert and do all that for nothing.
				continue;
			} 
			blockID := sectionPalette[sectionID]
			w.WriteLittleEndianUInt16(blockID)
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
