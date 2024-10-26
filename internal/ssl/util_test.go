package ssl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_realPath(t *testing.T) {
	tests := map[string]struct {
		path string
		want assert.ValueAssertionFunc
	}{
		"absolute path": {
			path: "/tmp/test",
			want: func(t assert.TestingT, i interface{}, _ ...interface{}) bool {
				return assert.Equal(t, "/tmp/test", i)
			},
		},
		"relative path with dot": {
			path: "./util_test.go",
			want: func(t assert.TestingT, i interface{}, _ ...interface{}) bool {
				v, ok := i.(string)
				return ok &&
					assert.Less(t, 1, len(v)) &&
					assert.NotEqual(t, '.', v[0]) &&
					assert.FileExists(t, v)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.want(t, realPath(tt.path))
		})
	}
}
