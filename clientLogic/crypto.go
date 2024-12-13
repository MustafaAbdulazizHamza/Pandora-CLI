package clientLogic

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// Encrypt data with the RSA public key
func encryptWithPublicKey(data, publicKeyPath string) (string, error) {
	// Read the public key file
	publicKeyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return "", fmt.Errorf("could not read public key file: %v", err)
	}

	// Decode the PEM-encoded public key
	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the public key using PKIX format (instead of PKCS1)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %v", err)
	}

	// Assert that the public key is of type *rsa.PublicKey
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("not an RSA public key")
	}

	// Encrypt the data using RSA OAEP
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, rsaPublicKey, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("encryption failed: %v", err)
	}

	// Return the encrypted data as a base64-encoded string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt the data with the RSA private key (handling PKCS8 format)
func decryptWithPrivateKey(encryptedData, privateKeyPath string) (string, error) {
	// Read the private key file
	privateKeyData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("could not read private key file: %v", err)
	}

	// Decode the PEM-encoded private key
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block containing private key")
	}

	// Try to parse the private key as PKCS8
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// Assert that the private key is of type *rsa.PrivateKey
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("not an RSA private key")
	}

	// Decode the base64-encoded encrypted data
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %v", err)
	}

	// Decrypt the data using RSA OAEP
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, rsaPrivateKey, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %v", err)
	}

	// Return the decrypted data as a string
	return string(plaintext), nil
}
