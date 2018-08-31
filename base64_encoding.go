package cryptu

import (
	"encoding/base64"
	"errors"
)

// Base64Encoding is an injectable wrapper around the
// encoding used by Base64.
type Base64Encoding interface {
	// Encoding returns the *Encoding value
	Encoding() *base64.Encoding
}

// base64Encoding is an implementation of Base64Encoding
type base64Encoding struct {
	value *base64.Encoding
}

// NewBase64Encoding returns a new instance of Base64Encoding from
// an existing encoding.
func NewBase64Encoding(encoding *base64.Encoding) (Base64Encoding, error) {
	if encoding == nil {
		return nil, errors.New("crypt:NewBase64Encoding encoding cannot be nil")
	}

	return &base64Encoding{encoding}, nil
}

func (b *base64Encoding) Encoding() *base64.Encoding {
	return b.value
}
