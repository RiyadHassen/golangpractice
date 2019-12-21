package main

import "fmt"

func encode(n int) (bool, string) {
	decimal := []int{1, 4, 5, 9, 10, 40,
		50, 90, 100, 400,
		500, 900, 1000}
	romrep := ""
	encoded := false
	roman := []string{
		"I", "IV", "V", "IX",
		"X", "XL", "L", "XC",
		"C", "CD", "D", "CM", "M"}
	for i := len(decimal); i > 0; i-- {
		if decimal[i] <= n {
			romrep += roman[i]
			n -= decimal[i]
		}
	}
	if roman != nil {
		return encoded, romrep
	} else {
		return encoded, romrep
	}

}

func main() {
	r, e := encode(1998)
	fmt.Println(r, e)
}