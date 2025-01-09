package palindromechecker

import "fmt"

func main() {
	fmt.Println(isPermutationPalindrome("Tact Coa"))
}

// this only checks if string can be re arranged into palindrome.

// string length could be odd or even.
// In a palindrome string with even length, every character must appear an even number of times.
// In a palindrome with an odd length, all but one character must appear an even number of times, with one character allowed to have an odd count.
func isPermutationPalindrome(s string) bool {
	var bitVector int32
	// small alphabet letters have large value compared to capital alphabet letters.
	for _, char := range s {
		// Convert to lowercase if it's an uppercase letter
		if char >= 'A' && char <= 'Z' {
			char += 32 // difference in ascii value for same alphabet is constant 32.
		}

		// Skip non-lowercase letters
		if char < 'a' || char > 'z' {
			continue
		}

		// Toggle the bit for this character
		bitPosition := char - 'a' // using this so that 32 bit vector can handle it between 0 to 26.
		bitVector ^= int32(1) << bitPosition
	}

	return bitVector == 0 || hasOnlyOneBitSet(bitVector)
}

func hasOnlyOneBitSet(n int32) bool {
	return n&(n-1) == 0
}

