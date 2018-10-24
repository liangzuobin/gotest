package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"0010000617747": 0,
		"0010008120458": 0,
		"0010011293082": 0,
		"0010010071411": 0,
	}

	for k, _ := range m {
		for i := 10000000; ; i++ {
			n := number(i)
			if k == n {
				m[k] = i
				fmt.Printf("%s = %d \n", k, i)
				break
			}
		}
	}

}

func number(id int) string {
	top, low := id/10000, id%10000
	return fmt.Sprintf("001%010d", low*1234+top)
}
