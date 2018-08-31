package cryptu_test

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/clavoie/cryptu"
	"github.com/clavoie/di"
)

func NewKey() (cryptu.Key, error)                 { return cryptu.NewStrKey("0dJFIeW64bTgtjTU") }
func NewEncoding() (cryptu.Base64Encoding, error) { return cryptu.NewBase64Encoding(base64.StdEncoding) }

var defs = []*di.Def{
	{NewKey, di.Singleton},
	{NewEncoding, di.Singleton},
}

func errFn(err *di.ErrResolve, w http.ResponseWriter, r *http.Request) {}

func ExampleBase64() {
	resolver, err := di.NewResolver(errFn, cryptu.NewDiDefs(), defs)

	if err != nil {
		log.Fatal(err)
	}

	resolver.Invoke(func(encoder cryptu.Base64) {
		secret := "my secret"

		encryptedValue, err := encoder.Encrypt(secret)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(encryptedValue == secret)

		decodedValue, err := base64.StdEncoding.DecodeString(encryptedValue)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(decodedValue) == secret)

		decryptedSecret, err := encoder.Decrypt(encryptedValue)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(decryptedSecret)
	})

	// Output: false
	// false
	// my secret
}
