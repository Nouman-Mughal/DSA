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

func bitwiseComplement(a int) int {
	return ^a
}
