package cryptu_test

import (
	"strings"
	"testing"

	"github.com/clavoie/cryptu"
)

func TestAes(t *testing.T) {
	newKey := func(s string) cryptu.Key {
		k, err := cryptu.NewStrKey(s)

		if err != nil {
			t.Fatal(err)
		}

		return k
	}

	key16 := newKey(strings.Repeat("a", 16))
	key24 := newKey(strings.Repeat("b", 24))
	key32 := newKey(strings.Repeat("c", 32))
	keyInvalid := newKey("invalid")

	t.Run("NewAes", func(t *testing.T) {
		if _, err := cryptu.NewAes(nil); err == nil {
			t.Fatal("Expecting error")
		}

		if _, err := cryptu.NewAes(keyInvalid); err == nil {
			t.Fatal("Expecting error", keyInvalid)
		}

		if _, err := cryptu.NewAes(key16); err != nil {
			t.Fatal(err)
		}

		if _, err := cryptu.NewAes(key24); err != nil {
			t.Fatal(err)
		}

		if _, err := cryptu.NewAes(key32); err != nil {
			t.Fatal(err)
		}
	})

	secret := "my secret value"

	t.Run("EncryptDecrypt", func(t *testing.T) {
		cipher, err := cryptu.NewAes(key24)

		if err != nil {
			t.Fatal(err)
		}

		encryptedSecret, err := cipher.Encrypt([]byte(secret))

		if err != nil {
			t.Fatal(err)
		}

		if string(encryptedSecret) == secret {
			t.Fatal("Not encrypted", string(encryptedSecret), secret)
		}

		decryptedSecret, err := cipher.Decrypt([]byte(encryptedSecret))

		if err != nil {
			t.Fatal(err)
		}

		if string(decryptedSecret) != secret {
			t.Fatal("Not decrypted", string(decryptedSecret), secret)
		}
	})

	t.Run("DecryptBadBlock", func(t *testing.T) {
		cipher, err := cryptu.NewAes(key24)

		if err != nil {
			t.Fatal(err)
		}

		notEncrypted := []byte{1, 2, 3}
		_, err = cipher.Decrypt(notEncrypted)

		if err == nil {
			t.Fatal("Was expecting err")
		}
	})
}
