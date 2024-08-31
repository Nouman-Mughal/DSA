package main

import "strings"

func main() {
	println(mergeAlternately("abc", "pqrst"))
}

func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder

	i, j := 0, 0

	for i < len(word1) && j < len(word2) { // this loop will run untill shorter one got traversed.
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[j])
		i++
		j++
	}
	// IN case of longer strings chars remain to be add in the string.

	for i < len(word1) {
		builder.WriteByte(word1[i])
		i++
	}

	for j < len(word2) {
		builder.WriteByte(word2[j])
		j++
	}

	return builder.String()
}
