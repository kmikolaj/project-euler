package main

import "fmt"

func triplet(sum int) (int, int, int) {
	for a := 1; a < sum/3; a++ {
		for b := a + 1; b < (sum-a+1)/2; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return 0, 0, 0
}

func main() {
	a, b, c := triplet(1000)
	fmt.Println("Result:", a*b*c)
}
