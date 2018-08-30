package cryptu

// EncryptToBase64 takes a plain text string, encrypts it using AES, and returns
// the base64 encoded version of the string.
func EncryptToBase64(cryptKey, text string) (string, error) {
	key, err := NewStrKey(cryptKey)

	if err != nil {
		return "", err
	}

	return NewBase64(key).Encrypt(text)
}

// DecryptFromBase64 takes a string encoded with EncryptToBase64
// unencodes it, decrypts it, and returns the plaintext string
func DecryptFromBase64(cryptKey, text string) (string, error) {
	key, err := NewStrKey(cryptKey)

	if err != nil {
		return "", err
	}

	return NewBase64(key).Decrypt(text)
}
