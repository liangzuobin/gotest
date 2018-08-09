package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type payload struct {
	ReferenceID string `json:"referenceID,omitempty"`
	UserID      uint64 `json:"userID,omitempty"`
	Price       uint32 `json:"price,omitempty"`
}

func main() {
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe("127.0.0.1:3337", nil))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s := r.PostFormValue("payload")
		log.Printf("payload = %v", s)
		var p payload
		if err := json.Unmarshal([]byte(s), &p); err != nil {
			log.Printf("unmarshal payload failed, err: %v, payload: %v", err, s)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		if time.Now().Unix()%2 == 0 {
			fmt.Fprintln(w, `{"status":"OK","message":"successful","data":{"balance": 100}}`)
			return
		}
		fmt.Fprintln(w, `{"status":"FAILED","message":"balance insufficient","data":{"balance": 10}}`)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
