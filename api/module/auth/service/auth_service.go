package service

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

var (
	clientSecret = os.Getenv("CLIENT_SECRET")
)

func AuthorizeClient(clientKey string, timestamp string, xSignature string) error {
	expectedSignature, err := createXSignature(clientKey, timestamp)
	if err != nil {
		return err
	}
	if expectedSignature != xSignature {
		return fmt.Errorf("unauthorized")
	}
	return nil
}

func createXSignature(clientKey string, timestamp string) (string, error) {
	stringToSign := fmt.Sprintf("%s|%s", clientKey, timestamp)

	key, err := loadPrivateKey()
	if err != nil {
		return "", err
	}

	h := sha256.New()
	h.Write([]byte(stringToSign))
	outputHashByte := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, outputHashByte)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	// cSecret := strings.Trim(clientSecret, "-----BEGIN RSA PRIVATE KEY-----")
	// cSecret = strings.Trim(cSecret, "-----END RSA PRIVATE KEY-----")
	// cSecret = strings.TrimSpace(cSecret)

	block, _ := pem.Decode([]byte(clientSecret))
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
