package main

import "fmt"

func fibonacci_even_sum(max_value int) int {
	left := 1
	right := 2
	sum := 0
	for right < max_value {
		if right%2 == 0 {
			sum += right
		}
		right = right + left
		left = right - left
	}
	return sum
}

func main() {
	fmt.Println("Result:", fibonacci_even_sum(4000000))
}
