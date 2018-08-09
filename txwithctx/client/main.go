package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unsafe"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}

func main() {

	// sendreq("txpanic") // will lock

	sendreq("txwithctxpanic") // seems releases the lock

	sendreq("tx")
}

func sendreq(path string) {
	timer := time.Now()
	defer func() {
		log.Printf("%s cost %d", path, time.Since(timer)/time.Millisecond)
	}()

	resp, err := client.Get("http://localhost:3336/" + path)
	if err != nil {
		log.Printf("get %s failed, err: %v", path, err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read %s response failed, err: %v", path, err)
		return
	}
	log.Printf("%s response = %s", path, *(*string)(unsafe.Pointer(&b)))
}
