package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("t.UnixNano() = %d, t.UTC().UnixNano() = %d \n", now.UnixNano(), now.UTC().UnixNano())
}
