package main

import (
	"fmt"
	"strings"
)

func mapstring(upper func(x string) string, lst []string) []string {
	maps := []string{}
	for i := 0; i < len(lst); i++ {
		maps = append(maps, upper(lst[i]))
	}
	return maps
}

func main() {
	lst := []string{"h", "l", "m"}

	f := func(x string) string {
		return strings.ToUpper(x)
	}
	fmt.Println(mapstring(f, lst))
}
