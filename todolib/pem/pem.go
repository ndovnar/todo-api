package pem

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func readKey(path string) (*pem.Block, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read key %w", err)
	}

	key, _ := pem.Decode(bytes)
	if key == nil {
		return nil, fmt.Errorf("failed to decode key")
	}

	return key, nil
}

func ReadPriviateKey(path string) (*rsa.PrivateKey, error) {
	pem, err := readKey(path)
	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(pem.Bytes)
	if err != nil {
		return nil, err
	}

	key := parsedKey.(*rsa.PrivateKey)

	return key, nil
}

func ReadPublicKey(path string) (*rsa.PublicKey, error) {
	pem, err := readKey(path)
	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParsePKIXPublicKey(pem.Bytes)
	if err != nil {
		return nil, err
	}

	key := parsedKey.(*rsa.PublicKey)

	return key, nil
}
