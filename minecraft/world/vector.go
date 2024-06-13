package world

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v Vector) Add(vec Vector) Vector {
	v.X += vec.X
	v.Y += vec.Y
	v.Z += vec.Z
	return v
}

func (v Vector) Subtract(vec Vector) Vector {
	v.X -= vec.X
	v.Y -= vec.Y
	v.Z -= vec.Z
	return v
}

func (v Vector) Multiply(vec Vector) Vector {
	v.X *= vec.X
	v.Y *= vec.Y
	v.Z *= vec.Z
	return v
}

func (v Vector) MultiplyAll(factor float64) Vector {
	v.X *= factor
	v.Y *= factor
	v.Z *= factor
	return v
}

func (v Vector) LengthSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vector) Normalize() Vector {
	vecLength := v.Length()
	v.X /= vecLength
	v.Y /= vecLength
	v.Z /= vecLength

	return v
}
