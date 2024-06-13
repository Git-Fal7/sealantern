package world

type Position struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Yaw   float32 `json:"yaw"`
	Pitch float32 `json:"pitch"`
}

func (p Position) IntX() int32 {
	return int32(p.X * 32)
}
func (p Position) IntY() int32 {
	return int32(p.Y * 32)
}
func (p Position) IntZ() int32 {
	return int32(p.Z * 32)
}
func (p Position) IntYaw() uint8 {
	return uint8((p.Yaw / 360) * 256)
}
func (p Position) IntPitch() uint8 {
	return uint8((p.Pitch / 360) * 256)
}

func (p Position) Add(pos Position) Position {
	p.X += pos.X
	p.Y += pos.Y
	p.Z += pos.Z
	return p
}

func (p Position) Subtract(pos Position) Position {
	p.X -= pos.X
	p.Y -= pos.Y
	p.Z -= pos.Z
	return p
}

func (p Position) ToVector() Vector {
	return Vector{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
	}
}

func (p Position) ToBlockPosition() BlockPosition {
	return BlockPosition{
		X: int(p.X),
		Y: int(p.Y),
		Z: int(p.Z),
	}
}

type BlockPosition struct {
	X int
	Y int
	Z int
}
