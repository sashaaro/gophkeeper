package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSecret_IsNew(t *testing.T) {
	type fields struct {
		ID     uuid.UUID
		Name   string
		UserID uuid.UUID
	}
	tests := map[string]struct {
		fields fields
		want   bool
	}{
		"No ID": {
			fields: fields{Name: "user secret", UserID: uuid.Must(uuid.NewV6())},
			want:   true,
		},
		"ID not nil": {
			fields: fields{Name: "user secret", UserID: uuid.Must(uuid.NewV6()), ID: uuid.Must(uuid.NewV6())},
			want:   false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := Secret{
				ID:     tt.fields.ID,
				Name:   tt.fields.Name,
				UserID: tt.fields.UserID,
			}
			assert.Equalf(t, tt.want, s.IsNew(), "IsNew()")
		})
	}
}
