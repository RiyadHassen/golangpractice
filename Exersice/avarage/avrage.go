package main

import "fmt"

func main() {
	value := []float64{3.0,3.0,3.0}
	sum := 0.0
	for i :=0; i <len(value); i++ {
		sum += value[i]
	}
	fmt.Println("the avarage of the value is ", sum/float64(len(value)))
}
