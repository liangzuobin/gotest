package main

import "github.com/go-playground/log"

func main() {
	m := make(map[string]int, 2)
	s := make([]string, 0, len(m))
	for k, v := range m {
		log.Printf("k = %s, v = %d", k, v)
	}
}
