package main

import "fmt"

func findMin(sl []int) int {
	min := sl[0]
	for i := 1; i < len(sl); i++ {
		if min > sl[i] {
			min = sl[i]
		}
	}
	return min
}

func main() {
	sl := []int{4, 11, 3, 2, 7}
	fmt.Println(findMin(sl))

}
