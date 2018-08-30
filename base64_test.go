package cryptu_test

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/clavoie/cryptu"
)

type InvalidKey []byte

func (ik InvalidKey) Bytes() []byte { return []byte(ik) }

func TestBase64(t *testing.T) {
	invalidKey := InvalidKey([]byte("invalid"))
	validKey, err := cryptu.NewStrKey(strings.Repeat("a", 16))

	if err != nil {
		t.Fatal(err)
	}

	secret := "my secret text"
	base64Secret := base64.URLEncoding.EncodeToString([]byte(secret))

	t.Run("EncryptBadKey", func(t *testing.T) {
		encoder := cryptu.NewBase64(invalidKey)
		_, err := encoder.Encrypt(secret)

		if err == nil {
			t.Fatal("Was expecting err")
		}
	})
	t.Run("Encrypt", func(t *testing.T) {
		encoder := cryptu.NewBase64(validKey)
		encrypted, err := encoder.Encrypt(secret)

		if err != nil {
			t.Fatal(err)
		}

		if encrypted == base64Secret {
			t.Fatal("Secret is not encrypted")
		}

		_, err = base64.StdEncoding.DecodeString(encrypted)
		if err != nil {
			t.Fatal("Not base64 encoded", encrypted, err)
		}
	})

	encoder := cryptu.NewBase64(validKey)
	encryptedSecret, err := encoder.Encrypt(secret)

	if err != nil {
		t.Fatal(err)
	}

	t.Run("DecryptBadKey", func(t *testing.T) {
		decoder := cryptu.NewBase64(invalidKey)
		_, err := decoder.Decrypt(encryptedSecret)

		if err == nil {
			t.Fatal("Was expecting error")
		}
	})
	t.Run("DecryptNotBase64", func(t *testing.T) {
		decoder := cryptu.NewBase64(validKey)
		_, err := decoder.Decrypt("bad text")

		if err == nil {
			t.Fatal("Was expecting error")
		}
	})
}
