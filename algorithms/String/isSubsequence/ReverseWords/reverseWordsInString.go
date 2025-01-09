package ReverseWords

import "strings"
func reverseWords(s string) string {
	
	words := strings.Fields(s)  // split the sting based on isSpace checker if space then split.
	
	for i, j := 0, len(words) - 1; i < j; i, j = i+1, j-1 {  
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")  // join the elements of array and add space between each.
}
