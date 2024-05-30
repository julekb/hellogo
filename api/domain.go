package api

import (
	"github.com/google/uuid"
)

type DomainModel struct {
	ID   uuid.UUID
	Name string `json:"name"`
}
