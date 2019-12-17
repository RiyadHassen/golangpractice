package main

import "fmt"

type stack2 struct {
	a1 [10]string
}

func (s *stack2) push(no string) {
	for i, val := range s.a1 {
		if i >= len(s.a1) && i >= 0 {
			return
		} else if val == "" {
			s.a1[i] = no
			return
		}
	}
}
func (s *stack2) pop() string {

	for i := len(s.a1) - 1; i >= 0; i-- {
		if s.a1[i] != "" {
			val := s.a1[i]
			s.a1[i] = ""
			return val
		}
	}
	return ""
}

func main() {
	s := stack2{}
	fmt.Println(s)
	s.push("m")
	s.push("l")
	s.push("k")
	fmt.Println(s)

}
