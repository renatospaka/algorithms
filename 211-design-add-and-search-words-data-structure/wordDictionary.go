package main

type TrieNode struct {
	children map[byte]*TrieNode
	Word     bool
}

type WordDictionary struct {
	root *TrieNode
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{
		root: &TrieNode{
			children: make(map[byte]*TrieNode),
		},
	}
}

// // why this??!!!
// func Constructor() WordDictionary {
// 	return WordDictionary{
// 		root: &TrieNode{
// 			children: make(map[byte]*TrieNode),
// 		},
// 	}
// }

func (w *WordDictionary) AddWord(word string) {
	current := w.root

	for i := 0; i < len(word); i++ {
		if _, ok := current.children[word[i]]; !ok {
			current.children[word[i]] = &TrieNode{
				children: make(map[byte]*TrieNode),
			}
		}
		current = current.children[word[i]]
	}
	current.Word = true
}

func (w *WordDictionary) Search(word string) bool {
	return dfs(0, w.root, word)
}

func dfs(start int, node *TrieNode, word string) bool {
	current := node

	for i := start; i < len(word); i++ {
		asc := word[i]
		if asc == '.' {
			for _, child := range current.children {
				if dfs(i+1, child, word) {
					return true
				}
			}
			return false
		} else {
			if _, ok := current.children[asc]; !ok {
				return false
			}
			current = current.children[asc]
		}
	}
	return current.Word
}
