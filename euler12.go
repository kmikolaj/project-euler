package main

import "fmt"
import "math"

func divisors(number int) []int {
	var divisors []int
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
			if number != i*i {
				divisors = append(divisors, number/i)
			}
		}
	}
	return divisors
}

func first_triangle_number_with_n_divisors(n int) int {
	number := 1
	triangle := 1
	div := 1
	for div < n {
		number++
		triangle += number
		div = len(divisors(triangle))
	}
	return triangle
}

func main() {
	fmt.Println("Result:", first_triangle_number_with_n_divisors(500))
}
