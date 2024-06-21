package schematic

import (
	"bytes"
	"errors"
	"os"

	"github.com/git-fal7/sealantern/minecraft/blocks"
	"github.com/git-fal7/sealantern/minecraft/world"
	"github.com/git-fal7/sealantern/minecraft/world/chunk"
	"github.com/git-fal7/sealantern/minecraft/world/map_world"

	"github.com/seebs/nbt"
)

func LoadSchematic(file string) (*map_world.Map, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	br := bytes.NewReader(dat)
	tag, _, err := nbt.LoadCompressed(br)
	if err != nil {
		return nil, err
	}

	if tag.Type() != nbt.TypeCompound {
		return nil, errors.New("unknown nbt type")
	}

	c := tag.(nbt.Compound)

	materials := c["Materials"].(nbt.String)
	if materials.String() != "Alpha" {
		return nil, errors.New("incompatible schematic version")
	}

	width := int(c["Width"].(nbt.Short))
	height := int(c["Height"].(nbt.Short))
	length := int(c["Length"].(nbt.Short))

	m := &map_world.Map{
		Spawn:     world.Position{X: 0, Y: 0, Z: 0},
		Dimension: world.OVERWORLD,
		ChunkMap:  make(map[chunk.ChunkKey]*chunk.Chunk),
	}

	blocksArray := []int8(c["Blocks"].(nbt.ByteArray))
	data := []int8(c["Data"].(nbt.ByteArray))
	var index int
	for x := 0; x < width; x++ {
		for z := 0; z < length; z++ {
			for y := 0; y < height; y++ {
				index = y*width*length + z*width + x
				bigId := blocks.GetLegacyBlockState(int(uint8(blocksArray[index])), int(uint8(data[index])))
				name := blocks.GetLegacyFromState(int(bigId)).Name
				m.SetBlock(x, y, z, name, false)
			}
		}
	}
	return m, nil
}
