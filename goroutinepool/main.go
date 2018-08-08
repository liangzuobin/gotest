package main

import (
	"log"
)

type job struct {
}

type worker struct {
	Name  string
	Pool  chan chan job
	Queue chan job
	quit  chan bool
}

func newWorker(name string, pool chan chan job) worker {
	return worker{
		Name:  name,
		Pool:  pool,
		Queue: make(chan job),
		quit:  make(chan bool),
	}
}

func (w worker) start() {
	go func() {
		for {
			log.Printf("register worker %s to pool.", w.Name)
			w.Pool <- w.Queue
			select {
			case <-w.Queue:
				log.Printf("worker %s get a job, and done.", w.Name)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w worker) stop() {
	go func() {
		w.quit <- true
	}()
}

var (
	queue = make(chan job, 10)
)

func main() {
	consumer()
	go producer()
}

func producer() {
	for {
		log.Println("send a job to queue")
		queue <- job{}
	}
}

func consumer() {
	p := make(chan chan job, 100)
	w := newWorker("Wall-E", p)
	w.start()
	for j := range queue {
		go func(j job) {
			q := <-p
			q <- j
		}(j)
	}
}
