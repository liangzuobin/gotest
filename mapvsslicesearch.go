package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const size = 128

	source := make([]int, 0, size)
	target := make([]int, 0, size/2)
	count := 0
	rand.Seed(99)
	for {
		i := rand.Int()
		source = append(source, i)
		count++
		if count%2 == 0 {
			target = append(target, i)
		}
		if count > size {
			break
		}
	}

	start := time.Now().UnixNano()
	foo(source, target)
	end := time.Now().UnixNano()

	fmt.Printf("foo start = %d, end = %d, cost = %d \n", start, end, end-start)

	start2 := time.Now().UnixNano()
	bar(source, target)
	end2 := time.Now().UnixNano()

	fmt.Printf("bar start = %d, end = %d, cost = %d \n", start2, end2, end2-start2)
}

func foo(source, target []int) []int {
	m := make(map[int]bool, len(source))
	for _, i := range source {
		m[i] = true
	}

	s := make([]int, 0, len(target))
	for _, i := range target {
		if m[i] {
			s = append(s, i)
		}
	}
	return s
}

func bar(source, target []int) []int {
	s := make([]int, 0, len(target))
	for _, i := range target {
		for _, j := range source {
			if i == j {
				s = append(s, i)
			}
		}
	}
	return s
}
