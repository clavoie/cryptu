package cryptu_test

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"

	"github.com/clavoie/cryptu"
)

type badSymmetric struct{}

func (bs *badSymmetric) Encrypt(v []byte) ([]byte, error) {
	return nil, errors.New("bad symmetric encrypt")
}
func (bs *badSymmetric) Decrypt(v []byte) ([]byte, error) {
	return nil, errors.New("bad symmetric decrypt")
}

func TestBase64(t *testing.T) {
	validKey, err := cryptu.NewStrKey(strings.Repeat("a", 16))

	if err != nil {
		t.Fatal(err)
	}

	cipher, err := cryptu.NewAes(validKey)

	if err != nil {
		t.Fatal(err)
	}

	secret := "my secret text"
	encoding, err := cryptu.NewBase64Encoding(base64.StdEncoding)

	if err != nil {
		t.Fatal(err)
	}

	base64Secret := encoding.Encoding().EncodeToString([]byte(secret))

	encoder := cryptu.NewBase64(cipher, encoding)

	t.Run("Encrypt", func(t *testing.T) {
		encrypted, err := encoder.Encrypt(secret)

		if err != nil {
			t.Fatal(err)
		}

		if encrypted == base64Secret || encrypted == secret {
			t.Fatal("Secret is not encrypted")
		}

		_, err = encoding.Encoding().DecodeString(encrypted)
		if err != nil {
			t.Fatal("Not base64 encoded", encrypted, err)
		}

		_, err = encoder.Encrypt("")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("EncryptFailure", func(t *testing.T) {
		encoder := cryptu.NewBase64(new(badSymmetric), encoding)

		val, err := encoder.Encrypt("test 123")
		if err == nil {
			t.Fatal("Was expecting err")
		}

		if val != "" {
			t.Fatal("Was expecting empty str", val)
		}
	})

	t.Run("DecryptNotBase64", func(t *testing.T) {
		_, err := encoder.Decrypt("bad text")

		if err == nil {
			t.Fatal("Was expecting error")
		}
	})
	t.Run("DecryptNotEncrypted", func(t *testing.T) {
		base64Str := encoding.Encoding().EncodeToString([]byte("not encrypted"))
		_, err := encoder.Decrypt(base64Str)

		if err == nil {
			t.Fatal("Was expecting error")
		}
	})
	t.Run("Decrypt", func(t *testing.T) {
		encryptedStr, err := encoder.Encrypt(secret)

		if err != nil {
			t.Fatal(err)
		}

		actualSecret, err := encoder.Decrypt(encryptedStr)

		if actualSecret != secret {
			t.Fatal(actualSecret, secret)
		}
	})
}
