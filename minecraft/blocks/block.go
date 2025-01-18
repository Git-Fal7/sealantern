package blocks

import (
	"github.com/git-fal7/sealantern/pkg/material"
)

type Block struct {
	Id   uint16
	Data uint8
}

func (b Block) GetMaterial() material.Material {
	return material.FindMaterialByID(b.Id)
}

func (b Block) GetFullID() int {
	return int(b.Id << 4) | int(b.Data & 0xF)
}
