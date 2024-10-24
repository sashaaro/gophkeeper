package ssl

import (
	"crypto/tls"
	"fmt"
)

type Config struct {
	CAPath         string
	PublicKeyPath  string
	PrivateKeyPath string
	certificate    *tls.Certificate
}

// Certificate - Вернёт публичный/приватный ключи, если они указаны в конфигурации, либо nil. Ключи в формате PEM
func (c *Config) Certificate() (*tls.Certificate, error) {
	if c.certificate == nil {
		if c.PublicKeyPath == "" || c.PrivateKeyPath == "" {
			return nil, nil
		}
		cert, err := tls.LoadX509KeyPair(c.path(c.PublicKeyPath), c.path(c.PrivateKeyPath))
		if err != nil {
			return nil, fmt.Errorf("load key pair fails: %w", err)
		}
		c.certificate = &cert
	}
	return c.certificate, nil
}
