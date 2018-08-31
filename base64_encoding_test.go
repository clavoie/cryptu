package cryptu_test

import (
	"encoding/base64"
	"testing"

	"github.com/clavoie/cryptu"
)

func TestBase64Encoding(t *testing.T) {
	_, err := cryptu.NewBase64Encoding(nil)

	if err == nil {
		t.Fatal("Was expecting err")
	}

	baseEncoding := base64.StdEncoding
	encoding, err := cryptu.NewBase64Encoding(baseEncoding)

	if err != nil {
		t.Fatal(err)
	}

	if encoding.Encoding() != baseEncoding {
		t.Fatal(encoding.Encoding(), baseEncoding)
	}
}
