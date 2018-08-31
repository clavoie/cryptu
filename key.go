package cryptu

import "errors"

// TODO: need *base64.Encoding struct - using 2 different encodings

// Key represents the symmetric encryption key used in
// operations by this package.
type Key interface {
	// Bytes returns the encryption key byte slice.
	Bytes() []byte
}

// key is an implementation of Key
type key []byte

// NewStrKey creates a new Key from a string.
func NewStrKey(str string) (Key, error) {
	return NewKey([]byte(str))
}

// NewKey returns a new Key from a byte slice.
func NewKey(val []byte) (Key, error) {
	if len(val) == 0 {
		return nil, errors.New("cryptu:NewKey cipher key must not be 0 length")
	}

	valInternal := make([]byte, len(val))
	copy(valInternal, val)

	return key(valInternal), nil
}

func (k key) Bytes() []byte { return []byte(k) }
