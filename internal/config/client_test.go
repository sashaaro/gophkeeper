package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		got := NewClient()
		assert.Equal(t, DefaultServerAddress, got.ServerAddr)
	})
	t.Run("Specify ServerAddress", func(t *testing.T) {
		want := "127.0.0.1:111"
		require.NoError(t, os.Setenv("SERVER_ADDR", want))
		got := NewClient()
		assert.Equal(t, want, got.ServerAddr)
	})
}
