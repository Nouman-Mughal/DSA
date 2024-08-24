package main

func main() {
	var a, b int
	a = 12
	b = 25

	println(b)
	//here we can test our bitwise operators.

	println(bitwiseComplement(a))

}

// bitwise AND Operator.
// The output of bitwise binary AND is 1 if the corresponding bits of two operands is 1.
// If either bit of an operand is 0, the result of corresponding bit is evaluated to 0.
func bitwiseAND(a int, b int) int {
	return a & b
}

// The output of binary bitwise OR is 1 if at least one corresponding bit of two operands is 1.
func bitwiseOR(a int, b int) int {
	return a | b
}

//The result of bitwise binary XOR operator is 1 if the corresponding bits of two operands are opposite.

func bitwiseXOR(a int, b int) int {
	return a ^ b
}

// Bitwise compliment operator is an unary operator (works on only one operand). It changes 1 to 0 and 0 to 1.
// Go doesnt have ~ operator like C or other languages but we can do it using xor operator.
// bitwise compliment of signed Int is -(N +1) and for signed int its just (N + 1)

func bitwiseComplement(a int) int {
	return ^a
}

/*
*
Right shift operator shifts all bits towards right by certain number of specified bits.
If we’re starting with a negative number (a binary number where the leftmost bit is a 1),all the empty spaces are filled with a 1.
If we’re starting with a positive number (where the leftmost bit, or most significant bit, is a 0),
We do all this to preserve the sign.
then all the empty spaces are filled with a 0.
*/
func bitwiseRightShift(a int, n int) int {

	// Right shift by N.
	// ! right shift turns unsigned int into signed int.
	return a >> n
}
