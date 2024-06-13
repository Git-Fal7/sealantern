package world

type Dimension uint8

const (
	NETHER    Dimension = 0xFF // -1
	OVERWORLD Dimension = 0x00 // 0
	END       Dimension = 0x01 // 1
)
