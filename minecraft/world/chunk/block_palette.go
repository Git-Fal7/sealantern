package chunk

type BlockPalette interface {
	GetId(id uint16) int
	GetContent() []uint16
}
