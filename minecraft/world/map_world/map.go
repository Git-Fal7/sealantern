package map_world

import (
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

func (m *Map) SetBlock(x, y, z int, typ string, overrideAir bool) {
	// MASSIVE MEMORY IMPROVEMENT - We dont need to redeclare already aired blocks
	// overrideAir is here when we have block changes (when player breaks the block)
	if typ == "minecraft:air" && !overrideAir {
		return
	}
	chunk := m.GetChunk(int32(x)/16, int32(z)/16)
	chunkSection := chunk.GetSection(y/16, m.Dimension == world.OVERWORLD)
	chunkSection.SetBlock(x%16, y%16, z%16, typ)
}

func (m *Map) GetBlock(x, y, z int) int {
	chunk := m.GetChunk(int32(x)/16, int32(z)/16)
	chunkSection := chunk.GetSection(y/16, m.Dimension == world.OVERWORLD)
	return blocks.BLOCK_REGISTRY.GetBlockId(chunkSection.Palette.RecoverName(int(chunkSection.Blocks[(y%16)<<8|(z%16)<<4|(x%16)])))
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
		data, sectionBitmask := c.ToData(skyLight)
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
			SectionBitMask: sectionBitmask,
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
