package config

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	got := NewClient()
	if got.ServerAddr == "" {
		t.Errorf("ServerAddr should have default value")
	}
}
