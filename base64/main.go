package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	if err := foo(); err != nil {
		panic(err)
	}
}

func foo() error {
	target := "lbLZDr2RDKGUju1GVAGJprVa3W-GLNuN07EvLNYnkw-JjrorPrclYw=="
	encoding := base64.NewEncoding("")
	b, err := encoding.DecodeString(target)
	if err != nil {
		return fmt.Errorf("decode string failed, err: %v", err)
	}
	fmt.Printf("%s", string(b))
	return nil
}
