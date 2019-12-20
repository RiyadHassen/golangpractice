package main

import "fmt"

func findMax(sl []int) int {
	max := sl[0]
	for i := 1; i < len(sl); i++ {
		if max < sl[i] {
			max = sl[i]
		}
	}
	return max
}

func main() {
	sl := []int{4, 11, 3, 2, 7}
	fmt.Println(findMax(sl))

}
