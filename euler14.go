package main

import "fmt"

func collatz_sequence_length(n int) int {
	seq := 1
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		seq++
	}
	return seq
}

func longest_collatz_sequence(limit int) int {
	seq := 1
	number := 1
	for i := 1; i < limit; i += 2 {
		count := collatz_sequence_length(i)
		if count > seq {
			seq = count
			number = i
		}
	}
	return number
}

func main() {
	fmt.Println("Result:", longest_collatz_sequence(1000000))
}
