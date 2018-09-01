# cryptu [![GoDoc](https://godoc.org/github.com/clavoie/cryptu?status.svg)](http://godoc.org/github.com/clavoie/cryptu) [![Build Status](https://travis-ci.org/clavoie/cryptu.svg?branch=master)](https://travis-ci.org/clavoie/cryptu) [![codecov](https://codecov.io/gh/clavoie/cryptu/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/cryptu) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/cryptu)](https://goreportcard.com/report/github.com/clavoie/cryptu)

Encryption wrappers for Go. Currently only base64 symmetric encryption wrappers, with hooks to be used by a dependency injection system.

## Encrypt To Base64

The package level functions `EncryptToBase64` and `DecryptFromBase64` use the `crypto/aes` package for encryption and `encoding/base64.StdEncoding` for the encoding. 

```go
// keys must be 16, 24, or 32 in length
key := "zqpf8VWyrUP9j1gC"
secret := "sensitive data"

encryptedValue, err := cryptu.EncryptToBase64(key, secret)

if err != nil {
  log.Fatal(err)
}

fmt.Println(encryptedValue)
// R+gUlOWekeVALBOntneoP7wQK2IOBiC3ddS+Rj2x

decodedValue, err := base64.StdEncoding.DecodeString(encryptedValue)

if err != nil {
  log.Fatal(err)
}

fmt.Println(string(decodedValue)) 
// G¦¦?¦@,¦¦w¦?¦+b ¦u?F=¦

decryptedSecret, err := cryptu.DecryptFromBase64(key, encryptedValue)

if err != nil {
  log.Fatal(err)
}

fmt.Println(decryptedSecret)
// sensitive data
```
A full example is available [here](https://godoc.org/github.com/clavoie/cryptu#example-EncryptToBase64)

## Dependency Injection

An interface is provided to wrap all top level package functions. This interface can be injected the into your code instead of calling the package functions directly. There are some [predefined dependency definitions](https://github.com/clavoie/cryptu/blob/master/di.go#L7) provided by the package. If you'd like to start using them you need only supply a definition for `cryptu.Key` and `cryptu.Base64Encoding`:

```go

func NewKey() (cryptu.Key, error)                 { return cryptu.NewStrKey(os.Getenv("MY_KEY")) }
func NewEncoding() (cryptu.Base64Encoding, error) { return cryptu.NewBase64Encoding(base64.StdEncoding) }

var defs = []*di.Def{
    {NewKey, di.Singleton},
    {NewEncoding, di.Singleton},
}

func MyHandler(encoder cryptu.Base64) {
  encryptedValue, err := encoder.Encrypt("some value")
  // if err
  
  myResponse := &MyResponse{Secret: encryptedValue}
  // write response
}

func main() {
  resolver, err := di.NewResolver(errFn, cryptu.NewDiDefs(), defs)
  // if err
  
  handler, err := resolver.HttpHandler(MyHandler)
  // if err
  
  http.HandleFunc("/foo", handler)
  // listen and serve
}
```

A full example is available [here](https://godoc.org/github.com/clavoie/cryptu#example-Base64)
