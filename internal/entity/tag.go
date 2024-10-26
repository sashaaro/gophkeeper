package entity

import "github.com/google/uuid"

type Tag struct {
	ID       uuid.UUID
	SecretID uuid.UUID
	Name     string
	Value    string
}
