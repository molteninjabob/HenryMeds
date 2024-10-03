package util

import "github.com/google/uuid"

func NewUUID() *uuid.UUID {
	id := uuid.New()
	return &id
}
