package main

import "log"

func main() {
	var f float32 = 18.30
	var i int32 = (int32)(f * 100)
	log.Println(i)
}
