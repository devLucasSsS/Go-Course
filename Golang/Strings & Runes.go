package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	//Strings
	sV1 := "A car"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	pl("Length: ", len(sV2))
	pl("Contains Another: ", strings.Contains(sV2, "Another"))
	pl("o index: ", strings.Index(sV2, "o"))
	pl("Replace: ", strings.Replace(sV2, "o", "0", -1))
	sv3 := "\nHola mundo \n"
	sv3 = strings.TrimSpace(sv3)
	pl("split: ", strings.Split("a-b-c-d", "-"))
	pl("Lower: ", strings.ToLower(sV2))
	pl("Upper: ", strings.ToUpper(sV2))
	pl("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	pl("suffix: ", strings.HasSuffix("tacocat", "cat"))
	//Runes
	rStr := "abcdefg"
	pl("Rune count: ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

}
