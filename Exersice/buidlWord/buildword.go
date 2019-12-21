package main

import "fmt"

func buildWord(word string,dict []string) int {
	//decleared to make copy of num
    res := 0
    done := false
	for i:=0;i<len(dict) ; i++ {
		sen := ""
		num := 0
		for j :=i; j<len(dict) ;j++  {
			sen += dict[j]
			num++
			if sen == word{
				// to break outer loop
				done = true
				break
			}
		}
		//to break out of the loop
		if done{
			res = num
			break
		}
	}

	return res

}

func main() {
	fmt.Println(buildWord("buildword",[]string{"buil","dwor","bu","ild","wo","rd"}))
	fmt.Println(buildWord("helloworld",[]string{"buil","dwor","bu","ild","wo","rd","hello","world"}))
	
}
