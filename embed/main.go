package main

import (
	"log"
)

type interFace interface {
	Do()
}

type impl struct {
	action string
}

func (i impl) Do() {
	log.Println(i.action)
}

type embed struct {
	i interFace
}

func (e embed) Do() {
	e.i.Do()
}

func main() {
	i := impl{"run"}
	e := embed{i}
	e.Do()
}
