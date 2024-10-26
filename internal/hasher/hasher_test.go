package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasher_IsEqual(t *testing.T) {
	r := &Hasher{}
	p, err := r.Hash("test")
	require.NoError(t, err)
	assert.NoError(t, r.IsEqual(p, "test"))
	assert.Error(t, r.IsEqual(p, "wrong"))
}
