package main

import "fmt"

func f(x,y int) (int,int){
	return y,x
}

func main() {
	fmt.Println(f(7,2))
	fmt.Println(f(2,7))
}
