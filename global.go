package cryptu

import "encoding/base64"

// newDefaultBase64 returns a new Base64 instance using AES as the
// cipher and encoding/base64.StdEncoding.
func newDefaultBase64(cryptKey string) (Base64, error) {
	key, err := NewStrKey(cryptKey)

	if err != nil {
		return nil, err
	}

	aes, err := NewAes(key)

	if err != nil {
		return nil, err
	}

	encoding, err := NewBase64Encoding(base64.StdEncoding)

	if err != nil {
		return nil, err
	}

	return NewBase64(aes, encoding), nil
}

// EncryptToBase64 takes a plain text string, encrypts it using AES, and returns
// the base64 encoded version of the string using encoding/base64.StdEncoding.
func EncryptToBase64(cryptKey, text string) (string, error) {
	encrypter, err := newDefaultBase64(cryptKey)

	if err != nil {
		return "", err
	}

	return encrypter.Encrypt(text)
}

// DecryptFromBase64 takes a string encoded with EncryptToBase64
// unencodes it, decrypts it, and returns the plaintext string using
// encoding/base64.StdEncoding.
func DecryptFromBase64(cryptKey, text string) (string, error) {
	decrypter, err := newDefaultBase64(cryptKey)

	if err != nil {
		return "", err
	}

	return decrypter.Decrypt(text)
}
