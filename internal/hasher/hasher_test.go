package hasher

import "testing"

func TestHasher_Hash(t *testing.T) {
	tests := map[string]struct {
		salt []byte
		str  string
		want string
	}{
		"with salt": {
			salt: []byte{1, 2, 3},
			str:  "pass",
			want: "8ff1af0b7c0c9cc280d49a2fdf8846ac6f54e0cc70d495d16ddf8f791bb18b4f",
		},
		"empty salt": {
			salt: []byte{},
			str:  "pass",
			want: "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := &Hasher{
				salt: tt.salt,
			}
			if got := r.Hash(tt.str); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
