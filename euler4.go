package main

import "fmt"

func power(number, power int) int {
	result := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for power > 0 {
			result *= number
			power--
		}
	}
	return result
}

func is_palindrome(number int) bool {
	var left []int
	position := 0
	for number > 0 {
		if power(10, position+1) <= number {
			position++
			left = append(left, number%10)
		} else if power(10, position) > number {
			position--
			if left[position] != number%10 {
				break
			}
		}
		number = number / 10
	}
	return number == 0
}

func largest_palindrome_product(left, right int) int {
	largest := 0
	for i := 1; i <= left; i++ {
		for j := 1; j <= right; j++ {
			if is_palindrome(i * j) {
				if i*j > largest {
					largest = i * j
				}
			}
		}
	}
	return largest
}

func main() {
	fmt.Println("Result:", largest_palindrome_product(999, 999))
}
