package main

import (
	"fmt"
	"strconv"
)

func main() {
	var(
		bit string
		bit32 uint32
		bit64 uint64
		result uint32
	)

	bit = "00000010100101000001111010011100"
	bit64, _ = strconv.ParseUint(bit, 2, 32)
	bit32 = uint32(bit64)
	result = reverseBitsBetterPerformance(bit32)
	fmt.Printf("The binary string %s when reversed is %d\n", bit, result)

	bit = "11111111111111111111111111111101"
	bit64, _ = strconv.ParseUint(bit, 2, 32)
	bit32 = uint32(bit64)
	result = reverseBitsBetterPerformance(bit32)
	fmt.Printf("The binary string %s when reversed is %d\n", bit, result)
}

func reverseBits(num uint32) uint32 {
	var result uint32 = 0

	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		result += (bit << (31 - i))
	}
	return result
}

func reverseBitsBetterPerformance(num uint32) uint32 {
	var result uint32
	var n = 31

	for n != -1 {
		last := num & 1
		last <<= n
		result += last
		num >>= 1
		n--
	}
	return result
}
