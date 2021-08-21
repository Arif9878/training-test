package main

import (
	"fmt"
	"sort"
	"strings"
)

func Anagram(ArrayString []string) {
	list := make(map[string][]string)

	for _, word := range ArrayString {
		key := SortString(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		for _, w := range words {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}

func SortString(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	ArrayString := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	Anagram(ArrayString)
}
