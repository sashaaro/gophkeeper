package server

import (
	"context"
	"testing"

	"github.com/sashaaro/gophkeeper/pkg/gophkeeper"
	"github.com/stretchr/testify/require"
)

func TestKeeperServer_Ping(t *testing.T) {
	s := NewKeeperServer()
	got, err := s.Ping(context.Background(), &gophkeeper.Empty{})
	require.NoError(t, err)
	require.Equal(t, &gophkeeper.Empty{}, got)
}
