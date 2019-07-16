
[author: ataklychev](https://github.com/ataklychev/levenshtein)

I just translated into package format

`Levenshtein distance` - (editorial distance, editing distance) - the minimum number of operations to insert one character, delete one character and replace one character with another, necessary to turn one line into another. Measured for two strings, widely used in information theory and computational linguistics.

[source wiki](https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D1%81%D1%81%D1%82%D0%BE%D1%8F%D0%BD%D0%B8%D0%B5_%D0%9B%D0%B5%D0%B2%D0%B5%D0%BD%D1%88%D1%82%D0%B5%D0%B9%D0%BD%D0%B0)

## Using

```go
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
	fmt.Printf("Levenshtein number for <%s> and <%s> is %d\n", word1, word2, lev.Distance(word1, word2)) // 19
	fmt.Printf("Levenshtein number for <%s> and <%s> is %d\n", word1, word3, lev.Distance(word1, word3)) // 13
	fmt.Printf("Levenshtein number for <%s> and <%s> is %d\n", word1, word4, lev.Distance(word1, word4)) // 23
}
```