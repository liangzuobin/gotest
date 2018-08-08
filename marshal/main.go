package main

import (
	"encoding/json"
	"log"
	"reflect"
	"runtime"

	"github.com/pkg/errors"
)

func main() {
	foo()
	bar()
}

func foo() {
	m := map[string]interface{}{
		"key3": 2.5,
		"key2": "value2",
		"key1": 1,
		"key4": nil,
		"key5": "",
		"key6": 0,
	}
	b, err := json.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("foo = %s", string(b))
}

type demo struct {
	Name    string  `json:"name"`
	Code    uint    `json:"code"`
	Balance float32 `json:"balance"`
}

func bar() {
	d := demo{
		Name:    "Joby",
		Code:    1,
		Balance: 25.1,
	}
	b, err := OrderedJSON(&d)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bar = %s", string(b))

	b, err = OrderedJSON(nil)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("bar = %s", string(b))

	b, err = OrderedJSON(&d)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bar = %s", string(b))
}

// OrderedJSON ...
func OrderedJSON(ptr interface{}) (b []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {

				// FIXME(liangzuobin) not so good
				err = errors.Errorf("OrderedJSON failed, with ptr = %+v", ptr)
			}
			if s, ok := r.(string); ok {
				err = errors.New(s)
			}
			err = r.(error)
		}
	}()
	v := reflect.ValueOf(ptr)
	e := reflect.Indirect(v)
	if num := e.NumField(); num > 0 {
		m := make(map[string]interface{}, num)
		for i := 0; i < num; i++ {
			m[e.Type().Field(i).Tag.Get("json")] = e.Field(i).Interface()
		}
		return json.Marshal(&m)
	}
	return []byte("{}"), nil
}
