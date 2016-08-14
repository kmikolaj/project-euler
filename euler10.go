package main

import "fmt"
import "math"

func sieve(limit int) []int {
	var numbers = make([]bool, limit)
	numbers[0] = true
	numbers[1] = true
	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		if numbers[i] == false {
			for j := 2 * i; j < limit; j += i {
				numbers[j] = true
			}
		}
	}
	var primes []int

	for i, value := range numbers {
		if !value {
			primes = append(primes, i)
		}
	}

	return primes
}

func sum(items []int) int {
	sum := 0
	for i := 0; i < len(items); i++ {
		sum += items[i]
	}
	return sum
}

func main() {
	fmt.Println("Result:", sum(sieve(2000000)))
}
