package chunk

type ChunkBlockPalette struct {
	Content []uint16
}

func (palette *ChunkBlockPalette) GetId(id uint16) int {
	for i, v := range palette.Content {
		if v == id {
			return i
		}
	}
	palette.Content = append(palette.Content, id)
	return len(palette.Content) - 1
}

func (palette *ChunkBlockPalette) GetContent() []uint16 {
	return palette.Content
}
