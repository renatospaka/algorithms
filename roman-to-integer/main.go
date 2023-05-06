package main

import (
	"fmt"
)

var roman map[string]int

func main() {
	roman = make(map[string]int)
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000

	fmt.Printf("%s is %d in integer numbers\n", "CMXCVIII", toInteger("CMXCVIII"))
}


// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
func toInteger(literals string) int {
	value := 0;
	
	for l := 0; l < len(literals); l++ {
		if l+1 < len(literals) && roman[string(literals[l])] < roman[string(literals[l+1])] {
			value -= roman[string(literals[l])]
		} else {
			value += roman[string(literals[l])]
		}
	}
	return value
}
