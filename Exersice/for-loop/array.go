package main

import "fmt"

func main() {
	loop := []int{}

	for i :=0; i < 10 ; i++  {
		loop = append(loop, i )
	}
	fmt.Println(loop)
}
