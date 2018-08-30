package cryptu

import "fmt"

// Key represents the symmetric encryption key used in
// operations by this package.
type Key interface {
	// Bytes returns the encryption key byte slice. The
	// slice must have a length of either 16, 24, or 32 .
	Bytes() []byte
}

// key is an implementation of Key
type key []byte

// NewStrKey creates a new Key from a string. The string
// must be 16, 24, or 32 bytes long, otherwise an error
// is returned.
func NewStrKey(str string) (Key, error) {
	return NewKey([]byte(str))
}

// NewKey returns a new Key from a byte slice. The slice must
// be 16, 24, or 32 bytes long.
func NewKey(val []byte) (Key, error) {
	switch len(val) {
	case 16:
	case 24:
	case 32:
		return key(val), nil
	}

	return nil, fmt.Errorf("cryptu:NewKey cipher key must be 16, 24, or 32 bytes long, but got a key that was %v bytes long", len(val))
}

func (k key) Bytes() []byte { return []byte(k) }
