package main

import (
	"errors"

	"log"
)

func main() {
	foo()
}

func foo() {
	i, err := bar()
	defer func() {
		if err != nil {
			log.Printf("i = %d, err = %v", i, err)
		}
	}()
	i = 2
	err = errors.New("err")
}

func bar() (int, error) {
	return 1, nil
}
