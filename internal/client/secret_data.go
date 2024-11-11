package client

import (
	"encoding/json"
	"github.com/sashaaro/gophkeeper/internal/entity"
)

type SecretType = byte

const (
	SecretTypeBinary          SecretType = 0x0
	SecretTypeText            SecretType = 0x1
	SecretTypeBankCredentials SecretType = 0x2
)

type SecretData []byte

func (s SecretData) GetType() SecretType {
	return SecretType(s[0]) // first byte determinate type of secret data
}

func (s SecretData) ToBinary() ([]byte, bool) {
	if SecretTypeBinary == s.GetType() {
		return s[1:], true
	}
	return nil, false
}
func (s SecretData) ToText() (string, bool) {
	if SecretTypeText == s.GetType() {
		return string(s[1:]), true
	}
	return "", false
}

func (s SecretData) ToBankCredentials() (entity.CreditCard, bool) {
	v := entity.CreditCard{}
	if SecretTypeBankCredentials == s.GetType() {
		_ = json.Unmarshal(s[1:], &v)
		return v, true
	}
	return v, false
}

func BytesToSecretData(s []byte) SecretData {
	return append([]byte{SecretTypeBinary}, s...)
}

func StringToText(s string) SecretData {
	return append([]byte{SecretTypeText}, []byte(s)...)
}

func CreditCardToSecretData(creditCard entity.CreditCard) SecretData {
	b, _ := json.Marshal(creditCard)
	return append([]byte{SecretTypeBankCredentials}, b...)
}
