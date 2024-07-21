//! Write an algorithm to reverse the String.

package main // to make this file standalone executables

import "fmt"

func main() {
	// fmt.Println(ReverseAciiString("Grüß Gott")) this will print bogus string
	fmt.Println(ReverseAsciiString("Hello"))

	fmt.Println(ReverseNonAsciiString("Grüß Gott"))
}

func ReverseAsciiString(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func ReverseNonAsciiString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
