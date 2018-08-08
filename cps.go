package main

import (
	"log"
	"time"
)

func sum(n int64) int64 {
	if n == 0 {
		return n
	}
	return n + sum(n-1)
}

func sum_tail(n, result int64) int64 {
	if n == 0 {
		return result
	}
	return sum_tail(n-1, result+n)
}

func sum_iter(n, result int64) int64 {
	for {
		if n == 0 {
			return result
		}
		n, result = n-1, result+n
	}
	return result
}

func sum_cps(n int64, f func(result int64) int64) int64 {
	if n == 0 {
		return f(n)
	}
	return sum_cps(n-1, func(result int64) int64 {
		return f(n + result)
	})
}

func foo(n int64) func(int64) int64 {
	return func(result int64) int64 {
		return n + result
	}
}

// 能算出数来，但是逻辑貌似是不对的
func sum_my_cps(n int64, f func(result int64) int64) int64 {
	if n == 0 {
		return f(n)
	}
	return sum_my_cps(n-1, foo(f(n)))
}

func main() {
	count := int64(1000)
	{
		start := time.Now().UnixNano()
		log.Println("sum = ", sum(count))
		end := time.Now().UnixNano()
		log.Println("sum cost ", end-start)
	}
	{
		start := time.Now().UnixNano()
		log.Println("sum_tail = ", sum_tail(count, 0))
		end := time.Now().UnixNano()
		log.Println("sum_tail cost ", end-start)
	}
	{
		start := time.Now().UnixNano()
		log.Println("sum_iter = ", sum_iter(count, 0))
		end := time.Now().UnixNano()
		log.Println("sum_iter cost ", end-start)
	}
	{
		start := time.Now().UnixNano()
		log.Println("sum_cps = ", sum_cps(count, func(i int64) int64 {
			return i
		}))
		end := time.Now().UnixNano()
		log.Println("sum_cps cost ", end-start)
	}
	{
		start := time.Now().UnixNano()
		log.Println("sum_my_cps = ", sum_my_cps(count, func(i int64) int64 {
			return i
		}))
		end := time.Now().UnixNano()
		log.Println("sum_my_cps cost ", end-start)
	}
}
