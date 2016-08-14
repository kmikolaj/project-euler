package main

import "fmt"

func sum35(limit int) int {
	sum := 0
	for i := 0; i < limit; i += 3 {
		sum = sum + i
	}
	for i := 0; i < limit; i += 5 {
		if i%3 != 0 {
			sum = sum + i
		}
	}
	return sum
}

func main() {
	fmt.Println("Result:", sum35(1000))
}
