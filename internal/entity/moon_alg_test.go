package entity_test

import (
	"testing"

	"github.com/sashaaro/gophkeeper/internal/entity"
)

func TestCheckMoonAlgorithm(t *testing.T) {
	tests := map[string]struct {
		b    string
		want bool
	}{
		"positive": {
			b:    "5062821234567892",
			want: true,
		},
		"positive2": {
			b:    "18",
			want: true,
		},
		"negative empty": {
			b:    "",
			want: false,
		},
		"negative one digit": {
			b:    "1",
			want: false,
		},
		"negative no number": {
			b:    "asdf123",
			want: false,
		},
		"negative": {
			b:    "5062821234567893",
			want: false,
		},
		"negative2": {
			b:    "5062821734567892",
			want: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := entity.CheckMoonAlgorithm(tt.b); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
