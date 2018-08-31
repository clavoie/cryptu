package cryptu

import (
	"encoding/base64"
)

// Base64 symmetrically encrypts / decrypts strings to and from
// base64 encoded values.
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
	cipher   Symmetric
	encoding *base64.Encoding
}

// NewBase64 returns a new implementation instance of Base64.
func NewBase64(cipher Symmetric, encoding Base64Encoding) Base64 {
	return &base64Impl{
		cipher:   cipher,
		encoding: encoding.Encoding(),
	}
}

func (b *base64Impl) Decrypt(value string) (string, error) {
	btext, err := b.encoding.DecodeString(value)

	if err != nil {
		return "", err
	}

	decryptedBytes, err := b.cipher.Decrypt([]byte(btext))

	if err != nil {
		return "", err
	}

	return string(decryptedBytes), nil
}

func (b *base64Impl) Encrypt(value string) (string, error) {
	encryptedValue, err := b.cipher.Encrypt([]byte(value))

	if err != nil {
		return "", err
	}

	return b.encoding.EncodeToString([]byte(encryptedValue)), nil
}
