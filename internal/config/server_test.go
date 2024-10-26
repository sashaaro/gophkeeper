package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		got := NewServer()
		assert.Equal(t, DefaultServerListen, got.Listen)
		assert.Equal(t, "", got.TLS.PublicKeyPath)
		assert.Equal(t, "", got.TLS.PrivateKeyPath)
		gotCert, err := got.TLS.Certificate()
		assert.NoError(t, err)
		assert.Nil(t, gotCert)
	})
}
