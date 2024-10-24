package server

import (
	"crypto/tls"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithTLS(t *testing.T) {
	o := GRPCServer{}
	cert := tls.Certificate{}
	require.Equal(t, 0, len(o.opts))
	WithTLS(&cert)(&o)
	require.Equal(t, 1, len(o.opts))
}
