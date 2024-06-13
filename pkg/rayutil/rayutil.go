package rayutil

import "github.com/git-fal7/sealantern/minecraft/world"

func GetVelocityRay(v world.Vector) world.Vector {
	velocityRay := v
	if velocityRay.LengthSquared() == 0 {
		velocityRay.X = 0
		velocityRay.Y = 1
		velocityRay.Z = 0
	} else {
		velocityRay = velocityRay.Normalize()
	}
	return velocityRay
}

func GetRayBetween(target world.Position, source world.Position) world.Vector {
	return target.Subtract(source).ToVector()
}
