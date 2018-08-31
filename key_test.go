package cryptu_test

import (
	"bytes"
	"testing"

	"github.com/clavoie/cryptu"
)

func TestKey(t *testing.T) {
	_, err := cryptu.NewStrKey("")

	if err == nil {
		t.Fatal("Was expecting err")
	}

	strKey := "my key"
	key, err := cryptu.NewStrKey(strKey)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal([]byte(strKey), key.Bytes()) == false {
		t.Fatal("Keys do not match", strKey, string(key.Bytes()))
	}
}
