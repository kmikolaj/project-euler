package main

import "fmt"

func max_prime_factor(value int) int {
	factor := 2
	for value > 1 {
		for value%factor != 0 {
			factor++
		}
		value /= factor
	}
	return factor
}

func main() {
	fmt.Println("Result:", max_prime_factor(600851475143))
}
