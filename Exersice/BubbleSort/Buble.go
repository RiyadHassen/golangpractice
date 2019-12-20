package main

import "fmt"

func main() {
	ls := []int{4, 3, 2, 6, 1}
	fmt.Println(bubleSort(ls))
}
func bubleSort(bub []int) []int {
	leng := len(bub)
	for i := 0; i < leng; i++ {
		fmt.Println(bub)
		for j := 1; j < leng; j++ {
			if bub[j-1] > bub[j] {
				bub[j-1], bub[j] = bub[j], bub[j-1]

			}
		}
		fmt.Println(bub)

	}
	return bub
}
