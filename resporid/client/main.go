package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type payload struct {
	ReferenceID string `json:"referenceID,omitempty"`
	UserID      uint64 `json:"userID,omitempty"`
	Price       uint32 `json:"price,omitempty"`
}

type svrresp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type goodResp struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

type badresp struct {
	ID string `json:"id"`
}

func main() {
	p := payload{
		ReferenceID: "reference_id",
		UserID:      1,
		Price:       100,
	}
	if err := consume(p); err != nil {
		log.Fatal(err)
	}
}

func consume(p payload) error {
	b, err := json.Marshal(&p)
	if err != nil {
		return fmt.Errorf("marshal payload failed, err: %v", err)
	}

	resp, err := http.PostForm("http://localhost:3336/consume", url.Values{"payload": []string{string(b)}})
	if err != nil {
		return fmt.Errorf("post form failed, err: %v", err)
	}
	defer resp.Body.Close()

	s := new(svrresp)
	var pending bool
	switch resp.StatusCode {
	case http.StatusAccepted:
		s.Data = new(badresp)
		pending = true
	case http.StatusOK:
		s.Data = new(goodResp)
	default:
		return fmt.Errorf("response failed, status code: %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return fmt.Errorf("decode resp failed, err: %v", err)
	}

	log.Printf("resp = %+v", s)

	if pending {
		r, ok := s.Data.(*badresp)
		if !ok {
			return errors.New("IMPOSSIBILE i see dead people")
		}
		query(r.ID)
	}

	return nil
}

func query(id string) {
	for {
		b, err := func() ([]byte, error) {
			resp, err := http.Get("http://localhost:3336/query?id=" + id)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			switch c := resp.StatusCode; c {
			case http.StatusAccepted:
				return nil, nil
			case http.StatusOK:
				b, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				}
				return b, nil
			default:
				return nil, fmt.Errorf("query failed, unexpected status: %s", resp.Status)
			}
		}()
		if err != nil {
			log.Fatal(err)
		}
		if len(b) > 0 {
			log.Printf("query response: %s", string(b))
			break
		}
		log.Println("another query will start in 1 second...")
		time.Sleep(1 * time.Second)
	}
}
