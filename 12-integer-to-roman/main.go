package main

import (
	"fmt"
	"strings"
)

type roman struct {
	symbol string
	value int
}
var roman2 []roman

func main() {
	roman2 = make([]roman, 13)
	roman2[0] = roman{value: 1, symbol: "I"}
	roman2[1] = roman{value: 4, symbol: "IV"}
	roman2[2] = roman{value: 5, symbol: "V"}
	roman2[3] = roman{value: 9, symbol: "IX"}
	roman2[4] = roman{value: 10, symbol: "X"}
	roman2[5] = roman{value: 40, symbol: "XL"}
	roman2[6] = roman{value: 50, symbol: "L"}
	roman2[7] = roman{value: 90, symbol: "XC"}
	roman2[8] = roman{value: 100, symbol: "C"}
	roman2[9] = roman{value: 400, symbol: "CD"}
	roman2[10] = roman{value: 500, symbol: "D"}
	roman2[11] = roman{value: 900, symbol: "CM"}
	roman2[12] = roman{value: 1000, symbol: "M"}

	fmt.Printf("%d is %s in roman notation\n", 998, toRoman(998))
	fmt.Printf("%d is %s in roman notation\n", 45, toRoman(45))
}


// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
func toRoman(number int) string {
	text := "";
	for l := len(roman2)-1; l > -1; l-- {
		// fmt.Printf("symbol: %v and value: %v and number/value: %d\n", roman2[l].symbol, roman2[l].value, number / roman2[l].value)

		if number / roman2[l].value > 0 {
			count := number / roman2[l].value
			text += strings.Repeat(roman2[l].symbol, count)
			number = number % roman2[l].value
		}
	}

	return text
}
