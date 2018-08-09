package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"unsafe"
)

func TestServer(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(handlePost))
	defer svr.Close()

	p := struct {
		ReferenceID string `json:"referenceID,omitempty"`
		UserID      uint64 `json:"userID,omitempty"`
		Price       uint32 `json:"price,omitempty"`
	}{
		ReferenceID: "reference_id",
		UserID:      1,
		Price:       1,
	}
	s, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	val := url.Values{
		"payload": {*(*string)(unsafe.Pointer(&s))},
	}

	fmt.Printf("body = %s \n", val.Encode())
	r := httptest.NewRequest(http.MethodPost, svr.URL, strings.NewReader(val.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	handlePost(w, r)

	if w.Code != http.StatusOK {
		t.Error("status not ok")
	}
}

func BenchmarkServer(b *testing.B) {
	b.StopTimer()
	svr := httptest.NewServer(http.HandlerFunc(handlePost))
	defer svr.Close()

	p := struct {
		ReferenceID string `json:"referenceID,omitempty"`
		UserID      uint64 `json:"userID,omitempty"`
		Price       uint32 `json:"price,omitempty"`
	}{
		ReferenceID: "",
		UserID:      1,
		Price:       1,
	}
	s, err := json.Marshal(p)
	if err != nil {
		b.Fatal(err)
	}
	val := url.Values{
		"payload": {*(*string)(unsafe.Pointer(&s))},
	}

	r := httptest.NewRequest(http.MethodPost, svr.URL, strings.NewReader(val.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		handlePost(w, r)
	}
}
