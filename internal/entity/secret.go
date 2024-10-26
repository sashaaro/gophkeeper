package entity

import "github.com/google/uuid"

type Secret struct {
	ID     uuid.UUID
	Name   string
	UserID uuid.UUID
}

func (s Secret) IsNew() bool {
	return s.ID == uuid.Nil
}
