package main

type Trie struct {
	children  [26]*Trie
	root      *Trie
	endOfWord bool
}

func Contructor() Trie {
	return Trie{
		children:  [26]*Trie{},
		root:      &Trie{},
		endOfWord: false,
	}
}

func (t *Trie) Insert(word string) {
	current := t.root

	for i := 0; i < len(word); i++ {
		index := word[i] - 'a' // only lowercase
		if current.children[index] == nil {
			current.children[index] = &Trie{}
		}
		current = current.children[index]
	}
	current.endOfWord = true
}

func (t *Trie) Search(word string) bool {
	current := t.root

	for i := 0; i < len(word); i++ {
		index := word[i] - 'a' // only lowercase
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.endOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root

	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a' // only lowercase
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return true
}
