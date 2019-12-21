package main

import "fmt"

func minnobuild(word string,dict []string) int  {
	frag :=""

	for _,val:=range word{
		frag = string(val)
		if stringin(frag,dict){

		}
	}
	return 0
}

func stringin(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
//check if the fragment iin the word
func stringinString(frag string,word string) bool{
   return false
}
func main() {
	fmt.Println(stringin("st",[]string{"st"}))
}
