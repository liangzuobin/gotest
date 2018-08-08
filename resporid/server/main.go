package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type payload struct {
	ReferenceID string `json:"referenceID,omitempty"`
	UserID      uint64 `json:"userID,omitempty"`
	Price       uint32 `json:"price,omitempty"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)

		// b, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	log.Printf("read request body failed, err: %v", err)
		// 	return
		// }
		// log.Printf("receive request, %v, %v", req.Method, string(b))

		switch req.Method {
		case http.MethodPost:

			var p payload
			s := req.FormValue("payload")
			if err := json.Unmarshal([]byte(s), &p); err != nil {
				log.Printf("unmarshal payload failed, err: %v, payload: %v", err, s)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Printf("payload = %v", p)
			w.WriteHeader(http.StatusOK)
			if time.Now().Unix()%2 == 0 {
				fmt.Fprintln(w, `{"status":"OK","message":"successful","data":{"balance": 100}}`)
				return
			}
			fmt.Fprintln(w, `{"status":"FAILED","message":"balance insufficient","data":{"balance": 10}}`)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:3337", nil))
}
