package cryptu

// Symmetric represents an implementation of a symmetric
// encryption algorithm.
type Symmetric interface {
	// Decrypt decrypts a byte slice encrypted with Encrypt() and
	// returns the unencrypted result.
	Decrypt([]byte) ([]byte, error)

	// Encrypt encrypts the byte slice using the underlying
	// algorithm, returning the encrypted bytes.
	Encrypt([]byte) ([]byte, error)
}
