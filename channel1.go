package main

import "fmt"

func main() {
	for {
		foo()
	}
}

func foo() {
	ch := make(chan int)
	defer close(ch)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover from panic")
			}
		}()
		ch <- 1
	}()
	if true {
		return
	}
	i := <-ch
	fmt.Println("int from channel = ", i)
}
