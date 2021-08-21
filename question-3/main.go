package main

import (
	"fmt"
	"strings"
)

func FindFirstStringInBracket(str string) string {
	IndexFirstBracketFound := strings.Index(str, "(")

	if IndexFirstBracketFound < 0 {
		return ""
	}

	IndexFirstBracketFound++

	IndexClosingBracketFound := strings.Index(str[IndexFirstBracketFound:], ")")

	if IndexClosingBracketFound < 0 {
		return ""
	}

	return str[IndexFirstBracketFound : IndexFirstBracketFound+IndexClosingBracketFound]
}

func main() {
	result := FindFirstStringInBracket("(stockbit)")
	fmt.Println(result)
}
