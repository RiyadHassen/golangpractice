package main

import "fmt"

func mapfunc(square func(x int) int, lst []int) []int {
	maps := []int{}
	for i := 0; i < len(lst); i++ {
		maps = append(maps, square(lst[i]))
	}
	return maps
}

func main() {
	lst := []int{2, 3, 4, 5}
	f := func(x int) int {
		return x * x
	}
	fmt.Println(mapfunc(f, lst))
}
