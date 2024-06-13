package profile

import "github.com/google/uuid"

type PlayerProfile struct {
	UUID       uuid.UUID
	Name       string
	Properties []Property
}
