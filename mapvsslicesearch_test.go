package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func baz(size int) (source, target []int) {
	source = make([]int, 0, size)
	target = make([]int, 0, size/2)
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
	return
}

func foo(source, target []int) []int {
	// m := make(map[int]bool, len(source))
	m := make(map[int]bool)
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

const (
	times = 10000
	size  = 64
)

func BenchmarkFoo(b *testing.B) {
	start := time.Now().UnixNano()
	for i := 0; i < times; i++ {
		source, target := baz(size)
		foo(source, target)
	}
	end := time.Now().UnixNano()
	fmt.Printf("\nfoo cost = %d ", (end-start)/times)
}

func BenchmarkBar(b *testing.B) {
	start := time.Now().UnixNano()
	for i := 0; i < times; i++ {
		source, target := baz(size)
		bar(source, target)
	}
	end := time.Now().UnixNano()
	fmt.Printf("\nbar cost = %d ", (end-start)/times)
}
