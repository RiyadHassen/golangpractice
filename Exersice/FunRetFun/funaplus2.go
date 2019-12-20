package main

func plus() func(a int) int {
	f := func(a int) int {
		return a + 2
	}
	return f
}
func main() {

}
