package cryptu

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// aes is an implementation of Symmetric
type aesImpl struct {
	key []byte
}

// NewAes returns a new instance of a Symmetric implemented using
// the crypto/aes package. The key must be 16, 24, or 32 bytes in
// length. Please see the crypto/aes package for details.
func NewAes(key Key) (Symmetric, error) {
	if key == nil {
		return nil, errors.New("cryptu:NewAes key is nil")
	}

	keyBytes := key.Bytes()
	switch len(keyBytes) {
	case 16:
	case 24:
	case 32:
	default:
		return nil, fmt.Errorf("cryptu:NewAes cipher key must be 16, 24, or 32 bytes long, but got a key that was %v bytes long", len(keyBytes))
	}

	return &aesImpl{key: keyBytes}, nil
}

func (a *aesImpl) Decrypt(encryptedValue []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)

	if err != nil {
		return nil, err
	}

	if len(encryptedValue) < aes.BlockSize {
		return nil, errors.New("cryptu:aes.Decrypt ciphertext too short")
	}

	btext := make([]byte, len(encryptedValue))
	copy(btext, encryptedValue)

	iv := btext[:aes.BlockSize]
	btext = btext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(btext, btext)

	return btext, nil

}

func (a *aesImpl) Encrypt(value []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)

	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(value))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], value)

	return ciphertext, nil
}
