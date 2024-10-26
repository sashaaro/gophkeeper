package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher struct {
	salt []byte
}

func NewHasher(salt string) *Hasher {
	return &Hasher{
		salt: []byte(salt),
	}
}

func (r *Hasher) Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	h.Write(r.salt)
	return hex.EncodeToString(h.Sum(nil))
}
