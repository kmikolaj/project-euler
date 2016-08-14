package main

import "fmt"
import "strconv"

func digitsum(n []int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}

func toArray(n int) []int {
	var a []int
	for n > 0 {
		a = append(a, n%10)
		n /= 10
	}
	return reverse(a)
}

func fromArray(n []int) string {
	str := ""
	for i := 0; i < len(n); i++ {
		str += strconv.Itoa(n[i])
	}
	return str
}

func reverse(n []int) []int {
	size := len(n)
	for i := 0; i < size/2; i++ {
		n[i], n[size-i-1] = n[size-i-1], n[i]
	}
	return n
}

func multiply(number []int, factor int) []int {
	r := 0
	i := 0
	n := reverse(number)
	for len(n) > i {
		m := n[i] * factor
		n[i] = m%10 + r%10
		r = m/10 + r/10
		i++
	}
	for r > 0 {
		n = append(n, r%10)
		r = r / 10
	}
	return reverse(n)
}

func power(n int, p int) []int {
	a := toArray(n)
	for i := 1; i < p; i++ {
		a = multiply(a, n)
	}
	return a
}

func main() {
	fmt.Println("Result:", digitsum(power(2, 1000)))
}
