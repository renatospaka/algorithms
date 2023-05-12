package main

import "fmt"

func main() {
	var(
		list = []int {}
		duplicate bool
	)

	list = []int{1, 2, 3, 1}
	duplicate = containDuplicates(list)
	fmt.Printf("The array %v contains duplicates? %t\n", list, duplicate)

	list = []int{1, 2, 3, 4}
	duplicate = containDuplicates(list)
	fmt.Printf("The array %v contains duplicates? %t\n", list, duplicate)

	list = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	duplicate = containDuplicates(list)
	fmt.Printf("The array %v contains duplicates? %t\n", list, duplicate)

	list = []int{3, 3}
	duplicate = containDuplicates(list)
	fmt.Printf("The array %v contains duplicates? %t\n", list, duplicate)
}

func containDuplicates(list []int) bool {
	set := make(map[int]int)
	fmt.Println("List:", list)
	for i := 0; i < len(list); i++ {
		_, ok := set[list[i]]
		fmt.Printf("ok: %t, key: %d, list[%d]: %d\n", ok, i, i, list[i])
		if ok {
			return true
		}

		set[list[i]] = list[i]
		fmt.Printf("set[i]: %d\n", set[list[i]])
	}

	return false
}
