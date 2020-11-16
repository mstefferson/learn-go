package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum_prev := 0
	temp := 0
	val_prev := 0
	return func() int {
		if sum_prev == 0 {
			sum_prev = 1
			return 0
		} else {
			temp = val_prev
			val_prev = sum_prev
			sum_prev = temp + sum_prev
			return val_prev
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
