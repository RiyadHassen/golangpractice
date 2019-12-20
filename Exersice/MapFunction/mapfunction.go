package main

//higher order function
//func square(x int) int  {
//	return x * x
//
//}

func mapfunc(square func(x int) int, lst []int) []int {
	maps := []int{}
	for i := 0; i < len(lst); i++ {
		maps = append(maps, square(i))
	}
	return maps
}

func main() {

}
