package main

import "fmt"

func main() {
	fmt.Println(fibonachi(6))
}

func fibonachi(x int) []int {
	fib := []int{}
	sum := 1
	fib = append(fib, sum)
	for i := 1; i < x; i++ {
		if i < 2 {
			fib = append(fib, i)
		} else {
			sum = fib[i-1] + fib[i-2]
			fib = append(fib, sum)
		}
	}
	return fib
}
