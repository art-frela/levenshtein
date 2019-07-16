package main

import (
	"fmt"

	lev "github.com/art-frela/levenshtein"
)

func main() {
	word1 := "Pirates of Caribbean sea"
	word2 := "Somali pirates"
	word3 := "Pirates of Caribbean sea, Black Pearl"
	word4 := "Пираты Карибского моря"
	fmt.Printf("Levenshtein number for <%s> and <%s> is %d\n", word1, word2, lev.Distance(word1, word2))
	fmt.Printf("Levenshtein number for <%s> and <%s> is %d\n", word1, word3, lev.Distance(word1, word3))
	fmt.Printf("Levenshtein number for <%s> and <%s> is %d\n", word1, word4, lev.Distance(word1, word4))
}
