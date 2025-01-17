package map_world

import (
	"bytes"
	"slices"

	"github.com/git-fal7/sealantern/minecraft/blocks"
	"github.com/git-fal7/sealantern/minecraft/player"
	"github.com/git-fal7/sealantern/minecraft/protocol/packet"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/chunk"
)

type Map struct {
	Spawn     world.Position
	Dimension world.Dimension
	ChunkMap  map[chunk.ChunkKey]*chunk.Chunk
}

func (m *Map) SetBlock(x, y, z int, blk blocks.Block) {
	chunk := m.GetChunk(int32(x)/16, int32(z)/16)
	chunkSection := chunk.GetSection(y/16, m.Dimension == world.OVERWORLD)
	chunkSection.SetBlock(x%16, y%16, z%16, uint16(blk.GetFullID()))
}

func (m *Map) SetBlockByID(x, y, z int, fullId uint16) {
	chunk := m.GetChunk(int32(x)/16, int32(z)/16)
	chunkSection := chunk.GetSection(y/16, m.Dimension == world.OVERWORLD)
	chunkSection.SetBlock(x%16, y%16, z%16, fullId)
}

func (m *Map) GetBlock(x, y, z int) blocks.Block {
	chunk := m.GetChunk(int32(x)/16, int32(z)/16)
	chunkSection := chunk.GetSection(y/16, m.Dimension == world.OVERWORLD)
	fullId := int(chunkSection.Palette.GetContent()[chunkSection.Blocks[(y%16)<<8|(z%16)<<4|(x%16)]])
	return blocks.Block{
		Id:   uint8(fullId >> 0x4),
		Data: uint8(fullId & 0xf),
	}
}

func (m *Map) GetChunk(x int32, z int32) *chunk.Chunk {
	key := chunk.ChunkKey{X: x, Z: z}
	c, ok := m.ChunkMap[key]
	if ok {
		return c
	}

	c = &chunk.Chunk{
		ChunkX: x,
		ChunkZ: z,
	}
	m.ChunkMap[key] = c
	return c
}

func (m *Map) SendChunksAroundPlayer(p player.IPlayer) (map[chunk.ChunkKey]bool, map[chunk.ChunkKey]bool) {
	prevChunks := make(map[chunk.ChunkKey]bool)
	for k, v := range p.KnownChunks() {
		prevChunks[k] = v
	}
	newChunksSlice := make([]chunk.ChunkKey, 0)
	newChunks := make(map[chunk.ChunkKey]bool)

	blockPos := p.Position().ToBlockPosition()
	centralX := int32(blockPos.X >> 4)
	centralZ := int32(blockPos.Z >> 4)

	radius := int32(min(8, 1+p.ClientSettings().ViewDistance))
	for x := (centralX - radius); x <= (centralX + radius); x++ {
		for z := (centralZ - radius); z <= (centralZ + radius); z++ {
			key := chunk.ChunkKey{X: x, Z: z}
			if _, ok := prevChunks[key]; ok {
				delete(prevChunks, key)
			} else {
				newChunksSlice = append(newChunksSlice, key)
				newChunks[key] = true
			}
		}
	}

	if len(newChunks) == 0 && len(prevChunks) == 0 {
		return nil, nil
	}

	slices.SortFunc(newChunksSlice, func(a chunk.ChunkKey, b chunk.ChunkKey) int {
		dx := 16*float64(a.X) + 8 - p.Position().X
		dz := 16*float64(a.Z) + 8 - p.Position().Z
		da := dx*dx + dz*dz
		dx = 16*float64(b.X) + 8 - p.Position().X
		dz = 16*float64(b.Z) + 8 - p.Position().Z
		db := dx*dx + dz*dz
		if da > db {
			return 1
		} else if da == db {
			return 0
		} else {
			return -1
		}
	})

	for key := range newChunks {
		p.KnownChunks()[key] = true
	}

	bulkSize := 6
	skyLight := m.Dimension == world.OVERWORLD
	packets := make([]packet.PacketPlayChunkData, 0)
	for _, key := range newChunksSlice {
		c := m.GetChunk(key.X, key.Z)
		data, bitmask := c.Data()
		if len(data) == 0 {
			c.Reload(skyLight)
			data, bitmask = c.Data()
		}
		messageSize := 10 + len(data)
		if bulkSize+messageSize > 0x1fffef {
			p.WritePacket(&packet.PacketPlayMapChunkBulk{
				Packets:  packets,
				Skylight: skyLight,
			})
			packets = make([]packet.PacketPlayChunkData, 0)
			bulkSize = 6
		}
		bulkSize += messageSize
		packets = append(packets, packet.PacketPlayChunkData{
			X:              c.ChunkX,
			Z:              c.ChunkZ,
			SectionBitMask: bitmask,
			Data:           data,
		})
	}
	// Leftovers
	if len(packets) != 0 {
		p.WritePacket(&packet.PacketPlayMapChunkBulk{
			Packets:  packets,
			Skylight: skyLight,
		})
	}
	// remove the old chunks from prevChunks
	for key := range prevChunks {
		p.WritePacket(&packet.PacketPlayChunkData{
			X:              key.X,
			Z:              key.Z,
			GroundUp:       true,
			SectionBitMask: 0,
			Data:           make([]byte, 0),
		})
		delete(p.KnownChunks(), key)
	}
	return newChunks, prevChunks
}

func (m *Map) Copy() *Map {
	var chunkMap map[chunk.ChunkKey]*chunk.Chunk = make(map[chunk.ChunkKey]*chunk.Chunk)
	for key, chunkValue := range m.ChunkMap {
		var sectionSlice [16]*chunk.ChunkSection
		for i, section := range chunkValue.Sections {
			if section == nil {
				continue
			}
			sectionPalette := section.Palette.(*chunk.ChunkBlockPalette)
			sectionBlocks := bytes.Clone(section.Blocks[:])
			sectionSkylight := bytes.Clone(section.SkyLight[:])
			sectionBlockLight := bytes.Clone(section.BlockLight[:])
			sectionSlice[i] = &chunk.ChunkSection{
				Palette: &chunk.ChunkBlockPalette{
					Content: sectionPalette.Content,
				},
				Blocks:     [4096]byte(sectionBlocks),
				SkyLight:   [2048]byte(sectionSkylight),
				BlockLight: [2048]byte(sectionBlockLight),
			}
		}
		chunkMap[key] = &chunk.Chunk{
			ChunkX:   chunkValue.ChunkX,
			ChunkZ:   chunkValue.ChunkZ,
			Sections: sectionSlice,
			Biomes:   [256]byte(bytes.Clone(chunkValue.Biomes[:])),
		}
	}
	return &Map{
		Spawn:     m.Spawn,
		Dimension: m.Dimension,
		ChunkMap:  chunkMap,
	}
}
