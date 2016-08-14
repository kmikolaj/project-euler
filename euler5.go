package main

import "fmt"

func prime_factors(number int) []int {
	var list []int

	factor := 2
	for number > 1 {
		if number%factor == 0 {
			number /= factor
			list = append(list, factor)
		} else {
			factor++
		}
	}

	return list
}

func insert(item int, items []int) []int {

	left := 0
	right := len(items) - 1

	for right >= left {
		mid := (left + right) / 2
		if items[mid] > item {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	items = append(items, 0)
	copy(items[left+1:], items[left:])
	items[left] = item

	return items
}

func union(left, right []int) []int {
	union := left
	j := 0

	for _, i := range right {
		for j < len(left) && left[j] < i {
			j++
		}

		if j == len(left) || left[j] > i {
			union = insert(i, union)
		}

		j++
	}

	return union
}

func mult(number int) int {
	var factors []int
	for i := 2; i <= number; i++ {
		pf := prime_factors(i)
		factors = union(factors, pf)
	}

	result := 1
	for _, i := range factors {
		result = result * i
	}

	return result
}

func main() {
	fmt.Println("Result:", mult(20))
}
