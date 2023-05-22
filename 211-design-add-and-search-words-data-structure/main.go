package main

import "fmt"

func main() {
	var (
		dict *WordDictionary
		word string
	)

	dict = NewWordDictionary() 

	word = "dad"
	dict.AddWord(word)
	// fmt.Printf("The word <%s> exists? %t\n", word, dict.Search(word))

	word = "mad"
	dict.AddWord(word)
	// fmt.Printf("The word <%s> exists? %t\n", word, dict.Search(word))

	word = "bad"
	dict.AddWord(word)
	fmt.Printf("The word <%s> exists? %t\n", word, dict.Search(word))

	word = "pad"
	fmt.Printf("The word <%s> exists? %t\n", word, dict.Search(word))

	word = ".ad"
	fmt.Printf("The word <%s> exists? %t\n", word, dict.Search(word))

	word = "b.."
	fmt.Printf("The word <%s> exists? %t\n", word, dict.Search(word))
}
