package main

import "fmt"

func main() {
	var (
		num1, num2 int
		sum        int
	)

	num1, num2 = 1, 2
	sum = getSum(num1, num2)
	fmt.Printf("The sum of %d + %d is %d\n", num1, num2, sum)

	num1, num2 = 2, 3
	sum = getSum(num1, num2)
	fmt.Printf("The sum of %d + %d is %d\n", num1, num2, sum)

	num1, num2 =-1, 3
	sum = getSum(num1, num2)
	fmt.Printf("The sum of %d + %d is %d\n", num1, num2, sum)
}

func getSum(a int, b int) int {
	for b != 0 {
		temp := (a & b) << 1
		a = (a ^ b)
		b = temp
	}
	return a
}
