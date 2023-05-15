package main

import "fmt"

func main() {
	var (
		text string
		valid bool
	)

	text = "()"
	valid = isValid(text)
	fmt.Printf("the sequence %s is valid? %t\n", text, valid)

	text = "()[]{}"
	valid = isValid(text)
	fmt.Printf("the sequence %s is valid? %t\n", text, valid)

	text = "(]"
	valid = isValid(text)
	fmt.Printf("the sequence %s is valid? %t\n", text, valid)

	text = "({})"
	valid = isValid(text)
	fmt.Printf("the sequence %s is valid? %t\n", text, valid)
	
	text = "[(]"
	valid = isValid(text)
	fmt.Printf("the sequence %s is valid? %t\n", text, valid)
	
	text = "[(])"
	valid = isValid(text)
	fmt.Printf("the sequence %s is valid? %t\n", text, valid)

	text = "({2})"
	valid = isValid(text)
	fmt.Printf("the sequence %s is valid? %t\n", text, valid)
}

func isValid(text string) bool {
	possible := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	mapping := make([]byte, 0)

	for _, char := range []byte(text) {
		fmt.Println(char, string(char))
		pair, ok := possible[char]
		if !ok {
			mapping = append(mapping, char)
			continue
		}

		if len(mapping) == 0 {
			return false
		}

		if mapping[len(mapping) - 1] != pair {
			return false
		}

		mapping = mapping[:len(mapping) - 1]
	}

	return len(mapping) == 0
}
