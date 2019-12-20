package main

import "fmt"

func plusX(x int) func(a int) int {
	f := func(a int) int {
		return a + x
	}
	return f
}
func main() {
	pluss := plusX(2)
	fmt.Println(pluss(2))

}
