package main

import "fmt"

func is_prime(number int) bool {
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	} else if number%2 == 0 {
		return false
	} else {
		for i := 3; i <= number/2; i += 2 {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}

func nth_prime(n int) int {
	count := 0
	number := 1
	for count < n {
		number++
		if is_prime(number) {
			count++
		}
	}
	return number
}

func main() {
	fmt.Println("Result:", nth_prime(10001))
}
