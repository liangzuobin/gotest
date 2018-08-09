package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

var (
	client = http.Client{
		Transport: http.DefaultTransport,
		Timeout:   3 * time.Second,
	}
	mockDB sync.Map
)

type svrresp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type badresp struct {
	ID string `json:"id"`
}

func consume(w http.ResponseWriter, req *http.Request) {
	key := strconv.FormatInt(time.Now().Unix(), 10) // mock unique id
	ch := make(chan []byte)
	go process(req, key, ch)

	select {
	case <-time.Tick(1 * time.Second):
		s := svrresp{
			Status:  "pending",
			Message: "请求处理中",
			Data:    &badresp{ID: key},
		}
		b, err := json.Marshal(&s)
		if err != nil {
			http.Error(w, "请稍后再试", http.StatusInternalServerError)
			return
		}

		// DB / Queue / Cache etc
		mockDB.Store(key, "pending")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, string(b))
	case b := <-ch:

		// DB / Queue / Cache etc
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, *(*string)(unsafe.Pointer(&b)))
	}
}

func process(req *http.Request, key string, ch chan []byte) {
	defer close(ch)

	// b, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	log.Printf("read request body failed, err: %v", err)
	// 	ch <- []byte(`{"status":"FAILED","message":"internal server error","data":{}}`)
	// 	return
	// }
	// log.Printf("request body = %s", string(b))

	// FIXME(liangzuobin)
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, req.Body); err != nil {
		log.Printf("copy request body failed, err: %v", err)
		ch <- []byte(`{"status":"FAILED","message":"internal server error","data":{}}`)
		return
	}

	// time.Sleep(2 * time.Second)

	r, err := http.NewRequest(req.Method, "http://localhost:3337", &buf)
	if err != nil {
		log.Printf("new request failed, err: %v", err)
		ch <- []byte(`{"status":"FAILED","message":"internal server error","data":{}}`)
		return
	}
	defer req.Body.Close()

	for n, v := range req.Header {
		r.Header.Set(n, v[0])
	}

	// fake server delay
	// time.Sleep(3 * time.Second)

	resp, err := client.Do(r)
	if err != nil {
		log.Printf("forward request failed, err: %v", err)
		ch <- []byte(`{"status":"FAILED","message":"internal server error","data":{}}`)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("forward request failed, err: %v\n", err)
		ch <- []byte(`{"status":"FAILED","message":"internal server error","data":{}}`)
		return
	}

	mockDB.Store(key, *(*string)(unsafe.Pointer(&b)))

	ch <- b
}

func query(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	if len(key) == 0 {
		http.Error(w, "no id", http.StatusBadRequest)
		return
	}
	val, ok := mockDB.Load(key)
	if !ok {
		http.NotFound(w, r)
		return
	}
	resp, ok := val.(string)
	if !ok {
		http.Error(w, "no id", http.StatusInternalServerError)
		return
	}

	if resp == "pending" {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, resp)
}

func main() {
	http.HandleFunc("/consume", consume)
	http.HandleFunc("/query", query)
	log.Fatal(http.ListenAndServe("127.0.0.1:3336", nil))
}
