package cryptu_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/clavoie/cryptu"
)

func TestKey(t *testing.T) {
	key16 := strings.Repeat("a", 16)
	key24 := strings.Repeat("b", 24)
	key32 := strings.Repeat("c", 32)

	t.Run("NewKey16", func(t *testing.T) {
		key, err := cryptu.NewKey([]byte(key16))

		if err != nil {
			t.Fatal(err)
		}

		if bytes.Equal(key.Bytes(), []byte(key16)) == false {
			t.Fatal(key.Bytes(), key16)
		}
	})
	t.Run("NewKey24", func(t *testing.T) {
		key, err := cryptu.NewKey([]byte(key24))

		if err != nil {
			t.Fatal(err)
		}

		if bytes.Equal(key.Bytes(), []byte(key24)) == false {
			t.Fatal(key.Bytes(), key24)
		}
	})
	t.Run("NewKey32", func(t *testing.T) {
		key, err := cryptu.NewKey([]byte(key32))

		if err != nil {
			t.Fatal(err)
		}

		if bytes.Equal(key.Bytes(), []byte(key32)) == false {
			t.Fatal(key.Bytes(), key32)
		}
	})
	t.Run("NewKeyInvalid", func(t *testing.T) {
		_, err := cryptu.NewKey([]byte(key16 + "A"))

		if err == nil {
			t.Fatal("Was expecting err")
		}

		_, err = cryptu.NewKey(nil)

		if err == nil {
			t.Fatal("Was expecting err")
		}
	})
	t.Run("NewStrKey", func(t *testing.T) {
		_, err := cryptu.NewStrKey("")

		if err == nil {
			t.Fatal("Was expecting err")
		}

		key, err := cryptu.NewStrKey(key16)

		if err != nil {
			t.Fatal(err)
		}

		if string(key.Bytes()) != key16 {
			t.Fatal(string(key.Bytes()), key16)
		}
	})
}
