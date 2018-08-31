package cryptu_test

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/clavoie/cryptu"
)

func ExampleEncryptToBase64() {
	// keys must be 16, 24, or 32 in length
	key := "zqpf8VWyrUP9j1gC"
	secret := "sensitive data"

	encryptedValue, err := cryptu.EncryptToBase64(key, secret)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(encryptedValue == secret)

	decodedValue, err := base64.StdEncoding.DecodeString(encryptedValue)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decodedValue) == secret)

	decryptedSecret, err := cryptu.DecryptFromBase64(key, encryptedValue)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decryptedSecret)

	// Output: false
	// false
	// sensitive data
}
