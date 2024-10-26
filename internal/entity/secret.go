package entity

import "github.com/google/uuid"

type SecretKind string

const (
	KindCredentials SecretKind = "credentials"
	KindCreditCard  SecretKind = "credit_card"
	KindText        SecretKind = "text"
	KindBinary      SecretKind = "binary"
)

type Secret struct {
	ID     uuid.UUID
	Name   string
	UserID uuid.UUID
	Kind   SecretKind
}

func (s Secret) IsNew() bool {
	return s.ID == uuid.Nil
}
