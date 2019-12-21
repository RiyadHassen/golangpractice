package main

import (
	"fmt"
)

func reverseParen(s string) string {
	revpar := ""
	rev := ""
	if checkpar(s) {
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				for j := i + 1; j < len(s); j++ {

					if s[j] == ')' {
						i = j
						break
					} else {
						if s[j] == '(' {

						} else {
							char := string(s[j])
							rev = rev + char
						}

					}
				}

				rev = reverse(rev)
				revpar += rev

			} else if s[i] == ')' {

			} else {
				if s[i] == '(' {
					revpar += ""
				} else {
					revpar += string(s[i])
				}
			}
		}
	}
	return revpar
}
func reverse(s string) string {
	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}

func main() {
	fmt.Println(checkpar("(bar)"))
	fmt.Println(checkpar("foo(bar(baz))blim"))
	fmt.Println(reverseParen("foo(bar(baz))blim"))
	fmt.Println(reverseParen("foo(bar)bar"))
}

func checkpar(s string) bool {
	close, open := false, false
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open = true
			for j := i + 1; j < len(s); j++ {
				if s[j] == ')' {
					close = true
					break
				}
			}
		}
	}
	if open && close {
		return true
	} else {
		return false
	}
}
