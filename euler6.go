package main

import "fmt"

func sum_of_squares(limit int) int {
	sum := limit * (limit + 1) * (2*limit + 1) / 6
	return sum
}

func square_of_sum(limit int) int {
	sum := (1 + limit) * limit / 2
	return sum * sum
}

func main() {
	fmt.Println("Result:", square_of_sum(100)-sum_of_squares(100))
}
