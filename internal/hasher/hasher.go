package hasher

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (r *Hasher) Hash(pwd string) (ret string, err error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h), nil
}

func (r *Hasher) IsEqual(hashed, pwd string) error {
	h, err := hex.DecodeString(hashed)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(h, []byte(pwd))
}
