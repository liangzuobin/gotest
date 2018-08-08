package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/url"
)

func main() {

	var (
		param1    = "value1"
		param2    = "value2"
		timestamp = "12345678"
		nonce     = "abcedfg"
		key       = "higklmn"
	)

	values := url.Values{}
	values.Add("nonce", nonce)
	values.Add("timestamp", timestamp)
	values.Add("param2", param2)
	values.Add("param1", param1)

	// to_sign
	str := values.Encode()
	log.Printf("str = %s", str)

	// sign
	sign := hexSHA256([]byte(str + key))
	log.Printf("sign = %s", sign)
}

func hexSHA256(bytes []byte) string {
	hasher := sha256.New()
	hasher.Write([]byte(bytes))
	return fmt.Sprintf("%X", hasher.Sum(nil))
}
