package pem

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func DecodePrivateKey(base64Key string) (*rsa.PrivateKey, error) {
	pem, err := decodeKey(base64Key)
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

func DecodePublicKey(base64Key string) (*rsa.PublicKey, error) {
	pem, err := decodeKey(base64Key)
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

func decodeKey(base64Key string) (*pem.Block, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 %w", err)
	}

	block, _ := pem.Decode(decodedKey)
	if block == nil {
		return nil, fmt.Errorf("failed to decode key %w", err)
	}

	return block, nil
}

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
