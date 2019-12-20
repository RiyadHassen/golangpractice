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
				}
			}
			if found {
				anaram = append(anaram, val)
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
	dict := []string{"madam cuire"}
	w := "radium cam"
	fmt.Println(FindAnagrams(dict, w))

}
