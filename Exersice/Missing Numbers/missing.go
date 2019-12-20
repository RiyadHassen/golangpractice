package main

import "fmt"

func main() {
	fmt.Println(findMissing([]int{4, 2, 3}))
}

func findMissing(intarry []int) []int {
	max := findMax(intarry)
	missings := []int{}

	for i := 1; i < max; i++ {
		if numin(i, intarry) {
			continue
		} else {
			missings = append(missings, i)
		}

	}
	return missings
}
func numin(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func findMax(sl []int) int {
	max := sl[0]
	for i := 1; i < len(sl); i++ {
		if max < sl[i] {
			max = sl[i]
		}
	}
	return max
}
