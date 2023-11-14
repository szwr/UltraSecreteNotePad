package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func EncryptString(str string, pass string) (string, error) {
	key := make([]byte, 32)
	copy(key[:len(pass)], pass)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("Aes: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("Cipher: %v", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	c := aesgcm.Seal([]byte(str)[:0], nonce, []byte(str), nil)

	return hex.EncodeToString(c), nil
}

func DecryptString(hexEncrypted string, pass string) (string, error) {
	key := make([]byte, 32)
	copy(key[:len(pass)], pass)

	c, err := hex.DecodeString(hexEncrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	plaintext, err := aesgcm.Open(c[:0], nonce, c, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
