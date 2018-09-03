package cryptu_test

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/clavoie/cryptu"
)

func TestGlobal(t *testing.T) {
	t.Run("InvalidKey", func(t *testing.T) {
		if _, err := cryptu.EncryptToBase64("", "hello"); err == nil {
			t.Fatal("Expecting error")
		}

		if _, err := cryptu.DecryptFromBase64("invalid key", "hello"); err == nil {
			t.Fatal("Expecting error")
		}
	})

	t.Run("EncodingFailure", func(t *testing.T) {
		encodingBackup := base64.StdEncoding
		defer func() {
			base64.StdEncoding = encodingBackup
		}()

		base64.StdEncoding = nil
		key := strings.Repeat("a", 16)
		secret := "myt secret"

		encryptedValue, err := cryptu.EncryptToBase64(key, secret)
		if err == nil {
			t.Fatal("Was expecting err")
		}

		if encryptedValue != "" {
			t.Fatal("Was expecting empty string")
		}
	})

	t.Run("EncryptDecrypt", func(t *testing.T) {
		key := strings.Repeat("a", 16)
		secret := "myt secret"

		encryptedValue, err := cryptu.EncryptToBase64(key, secret)

		if err != nil {
			t.Fatal(err)
		}

		decodedValue, err := base64.StdEncoding.DecodeString(encryptedValue)
		if err != nil {
			t.Fatal(err)
		}

		if encryptedValue == secret || string(decodedValue) == secret {
			t.Fatal("Not encrypted", encryptedValue, secret, string(decodedValue))
		}

		actualValue, err := cryptu.DecryptFromBase64(key, encryptedValue)

		if err != nil {
			t.Fatal(err)
		}

		if actualValue != secret {
			t.Fatal(actualValue, secret)
		}
	})
}
