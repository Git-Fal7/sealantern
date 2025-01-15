package blocks

import (
	"github.com/git-fal7/sealantern/pkg/material"
)

type Block struct {
	Id   uint8
	Data uint8
}

func (b Block) GetMaterial() material.Material {
	return material.FindMaterialByID(uint16(b.Id))
}

func (b Block) GetFullID() int {
	return int((b.Id << 4) | (b.Data & 0xF))
}
