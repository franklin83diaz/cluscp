package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateKey() (privateKeyPEM []byte, publicKeyPEM []byte, err error) {
	// Generate private key
	privateRsa, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	// Generate public key
	publicRsa := &privateRsa.PublicKey

	// convert to PEM
	privateKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateRsa),
	})

	publicKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicRsa),
	})

	return
}
