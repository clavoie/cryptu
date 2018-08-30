package cryptu

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Base64 symmetrically encrypts / decrypts strings to and from
// base64 encoded values using AES.
type Base64 interface {
	// Decrypt decrypts a base64 encoded value and returns the decrypted
	// value.
	Decrypt(string) (string, error)

	// Encrypt encrypts a raw value and returns the encrypted base64 encoded
	// value.
	Encrypt(string) (string, error)
}

// base64Impl is an implementation of the Base64 interface.
type base64Impl struct {
	key Key
}

// NewBase64 returns a new implementation instance of Base64.
func NewBase64(key Key) Base64 {
	return &base64Impl{
		key: key,
	}
}

func (b *base64Impl) Decrypt(value string) (string, error) {
	block, err := aes.NewCipher(b.key.Bytes())

	if err != nil {
		return "", err
	}

	btext, err := base64.URLEncoding.DecodeString(value)

	if err != nil {
		return "", err
	}

	if len(btext) < aes.BlockSize {
		return "", errors.New("cryptu: ciphertext too short")
	}

	iv := btext[:aes.BlockSize]
	btext = btext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(btext, btext)

	return string(btext), nil
}

func (b *base64Impl) Encrypt(encryptedValue string) (string, error) {
	block, err := aes.NewCipher(b.key.Bytes())

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(encryptedValue))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(encryptedValue))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}
