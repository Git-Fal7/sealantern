package types

type BlockFace uint8

const (
	FaceDown BlockFace = iota
	FaceUp
	FaceNorth
	FaceSouth
	FaceWest
	FaceEast
)
