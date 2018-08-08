package main

import (
	"encoding/json"
	"fmt"
)

// Signal ...
type Signal struct {
	Name     string `json:"-"`
	Priority int    `json:"-"`
	Message  string `json:"-"`
	Location string `json:"-"`
	RealSignal
}

// RealSignal ...
type RealSignal struct {
	Name     string
	Priority int
	Message  string
	Location string
}

func main() {
	s := Signal{
		Name:     "fakeName",
		Priority: 0,
		Message:  "fakeMessage",
		Location: "fakeLocation",
		RealSignal: RealSignal{
			Name:     "realName",
			Priority: 10,
			Message:  "realMessage",
			Location: "realLocation",
		},
	}
	b, _ := json.Marshal(&s)
	receiveSignal(b)
}
func receiveSignal(data []byte) {
	if string(data) != "" {
		signal := &struct {
			Name     string
			Priority int
			Message  string
			Location string
		}{}
		json.Unmarshal(data, &signal)
		fmt.Println("\nRECEIVED AT THE BUREAU")
		fmt.Printf("signal = %+v \n", signal)
	}
}
