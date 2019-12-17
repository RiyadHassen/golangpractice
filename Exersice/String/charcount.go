package main

import "fmt"

func main() {
	stat := "asSASA ddd dsjkdsjs dk"
	for i := 0; i < len(stat) ;i++ {

	}

	for indx,runevalue := range stat{
		fmt.Printf("%#U start at byte postion %s\n",runevalue,indx)
	}

	
}
