package main

import "fmt"

func FindAnagrams(dictionary []string, word string) []string {
	anaram := []string{}
	found := false
	for _, val := range dictionary {
		for i := 0; i < len(val); i++ {
			for j := 0; j < len(word); j++ {
				if val[i] == word[j] {

					found = true
					break
				} else {
					found = false
					continue
				}
			}
		}

	}
	if found {
		return anaram
	} else {
		return nil
	}
}

func main() {
	dict := []string{"madam cuir,madam curi"}
	w := "radium came"
	fmt.Println(FindAnagrams(dict, w))

}
