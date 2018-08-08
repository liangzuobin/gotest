package main

import "log"

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func main() {
	list := []int{1, 2, 4}
	out := gen(list...)
	for i := 0; i < len(list); i++ {
		log.Println(<-out)
	}
}
