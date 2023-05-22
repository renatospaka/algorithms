package main

import "fmt"

func main() {
	var (
		trie Trie
		word string
		prefix string
	)

	trie = Contructor() // it would be nicier if using NewTrie instead of Constructor

	word = "apple"
	trie.Insert(word)
	fmt.Printf("The word <%s> exists? %t\n", word, trie.Search(word))

	prefix = "app"
	fmt.Printf("The word <%s> exists? %t\n", prefix, trie.Search(prefix))
	fmt.Printf("The prefix <%s> exists? %t\n", prefix, trie.StartsWith(prefix))

	word = prefix
	trie.Insert(word)
	fmt.Printf("The word <%s> exists? %t\n", word, trie.Search(word))
}
