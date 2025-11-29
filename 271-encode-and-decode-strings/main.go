package main

import "fmt"

func main() {
	test([]string{"leet", "code"})
	test([]string{"leet", "code", "love", "you"})
}

func test(list []string) {
	encoded := encodes(list)
	fmt.Printf("Encoded result of %v is %v", list, encoded)
	decoded := decodes(encoded)
	fmt.Printf(", decoded result is %v\n", decoded)
}

func encodes(list []string) string {
	result := ""
	for _, str := range list {
		result += fmt.Sprintf("%d#%s", len(str), str)
	}
	return result
}

func decodes(encoded string) []string {
	result := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length := 0
		fmt.Sscanf(encoded[i:j], "%d", &length)
		str := encoded[j+1 : j+1+length]
		result = append(result, str)
		i = j + 1 + length
	}
	return result
}
