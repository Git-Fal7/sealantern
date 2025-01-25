package entity

type LivingEntity interface {
	Entity
	Health() float32
	SetHealth(health float32)
}