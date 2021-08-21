package main

import (
	"fmt"
	"runtime"
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

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	ArrayString := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	go Anagram(ArrayString)

	var input string
	fmt.Scanln(&input)

	// go print(5, "halo")
	// print(5, "apa kabar")

	// var input string
	// fmt.Scanln(&input)

	// var list = make(map[string](chan string))

	// list := make(map[string][]string)

	// for _, word := range ArrayString {
	// 	key := SortString(word)
	// 	list[key] = append(list[key], word)
	// }

	// var baik = make(chan []string)
	// for _, words := range list {
	// 	go func(word []string) {
	// 		fmt.Println(word)
	// 		// list := make(map[string][]string)
	// 		// key := SortString(word)
	// 		// list[key] = append(list[key], word)
	// 		baik <- word
	// 	}(words)
	// }

	// go sayHelloTo("john wick")
	// go sayHelloTo("ethan hunt")
	// go sayHelloTo("jason bourne")

	// var message1 = <-messages
	// fmt.Println(baik)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(baik)
	// }
	// for words := range listCOk {
	// 	for _, w := range words {
	// 		fmt.Print(w, " ")
	// 	}
	// 	fmt.Println()
	// }
	// var message2 = <-messages
	// fmt.Println(message2)

	// var message3 = <-messages
	// fmt.Println(message3)
}
