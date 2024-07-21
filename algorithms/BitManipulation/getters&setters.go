package main

import "fmt"

func main() {
	fmt.Println(getBit(5, 1))
}

// get bit.
func getBit(num int, i int) bool {
	return ((num & (1 << i)) != 0)
}

// set bit.
func setBit(num int, i int) int {
	return num | (1 << i)
}

// clear bit.
func clearBit(num int, i int) int {
	mask := ^(1 << i) // go does not have negation but go uses bitwise XOR operator ^ for integer negation.
	return num & mask
}

func updateBit(num int, i int, bitValue bool) int {
	var value int
	if bitValue {
		value = 1
	} else {
		value = 0
	}
	return clearBit(num, i) | (value << i)

}
