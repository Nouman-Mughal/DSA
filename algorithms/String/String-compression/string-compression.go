package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(CompressString("aabbccccccdddddddddd"))
}

func CompressString(input string) string {

	if len(input) == 0 {
		return " "
	}

	var builder strings.Builder

	currentChar := input[0]
	charCount := 1

	for i := 1; i < len(input); i++ {
		if input[i] == currentChar {
			charCount++
		} else {
			writeCharAndCount(&builder, currentChar, charCount)
			currentChar = input[i]
			charCount = 1
		}
	}
	// we have to write this funtion outside of loop because loop ends before writing in the builder.
	writeCharAndCount(&builder, currentChar, charCount)

	return builder.String()

}
func writeCharAndCount(builder *strings.Builder, char byte, count int) {
	builder.WriteByte(char)
	if count > 1 {
		builder.WriteString(strconv.Itoa(count))
	}
}
