package ssl

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc/credentials"
)

type ServerConfig struct {
	PublicKeyPath  string
	PrivateKeyPath string
	certificate    *tls.Certificate
}

type ClientConfig struct {
	PublicKeyPath string
	certificate   credentials.TransportCredentials
}

// Certificate - Вернёт публичный/приватный ключи, если они указаны в конфигурации, либо nil. Ключи в формате PEM
func (c *ServerConfig) Certificate() (*tls.Certificate, error) {
	if c.certificate == nil {
		if c.PublicKeyPath == "" || c.PrivateKeyPath == "" {
			return nil, nil
		}
		cert, err := tls.LoadX509KeyPair(realPath(c.PublicKeyPath), realPath(c.PrivateKeyPath))
		if err != nil {
			return nil, fmt.Errorf("load key pair fails: %w", err)
		}
		c.certificate = &cert
	}
	return c.certificate, nil
}

func (c *ClientConfig) Certificate() (credentials.TransportCredentials, error) {
	if c.certificate == nil {
		if c.PublicKeyPath == "" {
			return nil, nil
		}
		cred, err := credentials.NewClientTLSFromFile(realPath(c.PublicKeyPath), "x.test.example.com")
		if err != nil {
			return nil, fmt.Errorf("load client tls fails: %w", err)
		}
		c.certificate = cred
	}
	return c.certificate, nil
}
