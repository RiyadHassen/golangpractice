package main

import (
	"fmt"
)

type stack1 struct {
	a1 [10]int
}

func (s *stack1) push(no int) {
	for i, val := range s.a1 {
		if i >= len(s.a1) && i >= 0 {
			return
		} else if val == 0 {
			s.a1[i] = no
			return
		}
	}
}
func (s *stack1) pop() int {

	for i := len(s.a1) - 1; i >= 0; i-- {
		if s.a1[i] != 0 {
			val := s.a1[i]
			s.a1[i] = 0
			return val
		}
	}
	return 0
}

func main() {

	s := stack1{}
	s.push(2)
	fmt.Println(s)
	s.push(4)
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s)
}
