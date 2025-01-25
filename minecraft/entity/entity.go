package entity

import "github.com/google/uuid"

type Entity interface {
	ID() int32
	UUID() uuid.UUID
}