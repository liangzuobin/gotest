package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	// apiToken := "a3ee6ea5-7e66-4d55-94bd-ab8692fe8111"
	// currentTimestamp := "1492571041"

	apiToken := "abcdefg"
	currentTimestamp := "12345678"
	signature := fmt.Sprintf("%x", sha1.Sum([]byte(apiToken+currentTimestamp)))
	fmt.Printf("signature = %s \n", signature)
}
