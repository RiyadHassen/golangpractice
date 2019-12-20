package main

import "fmt"

func main() {
	fmt.Println(findMissing2([]int{4, 2, 3}, 10))
}

func findMissing2(intarry []int, max int) []int {
	missings := []int{}

	for i := 1; i < max; i++ {
		if numcheck(i, intarry) {
			continue
		} else {
			missings = append(missings, i)
		}

	}
	return missings
}
func numcheck(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
