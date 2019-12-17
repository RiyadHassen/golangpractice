package main

import (
	"fmt"
)

func main() {
	s := "foobar"


	for i := len(s) -1;i >= 0; i--  {
		fmt.Printf("%s",string(s[i]))
	}
}
