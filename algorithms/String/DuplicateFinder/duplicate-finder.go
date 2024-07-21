// Here we are going to use bit manipulation technique to determine the duplicates in string
package main

func main() {
	println(hasDuplicates("hello"))                       // true
	println(hasDuplicates("world"))                       // false
	println(hasDuplicates("abcdefghijklmnopqrstuvwxyz"))  // false
	println(hasDuplicates("abcdefghijklmnopqrstuvwxyza")) // true

}

// this algorithm is only for ASCII string.So first asks the interviewer is String ASCII or Unicode.
// what are ASCII strings:
/*ASCII Strings:

  Character Set: ASCII (American Standard Code for Information Interchange) defines 128 characters.
  Encoding: Each character is represented by a 7-bit integer (0-127).
  Range: Includes basic Latin letters (A-Z, a-z), digits (0-9), and some punctuation marks.
  Size: Each character occupies 1 byte (8 bits, with the highest bit unused).
  Limitations: Cannot represent characters from most non-English languages or special symbols. */

/*
 	Unicode Strings:

  Character Set: Unicode defines over 140,000 characters, covering almost all of the world's writing systems.
  Encoding: Several encoding schemes exist, with UTF-8 being the most common:

  UTF-8: Variable-length encoding (1 to 4 bytes per character)
  UTF-16: Variable-length encoding (2 or 4 bytes per character)
  UTF-32: Fixed-length encoding (4 bytes per character)


  Range: Includes characters from virtually all languages, plus symbols, emojis, and more.
  Size: Varies based on the encoding used and the specific characters.
  Advantages: Can represent almost any character or symbol from any language.
*/

func hasDuplicates(s string) bool {
	var checker uint64 = 0
	// Each Ascii String Can have upto 128 unique characters
	if len(s) > 128 {
		return true
	}

	for i := 0; i < len(s); i++ {
		val := s[i] - 'a'
		/**
				Single Quotes in Go:

		In Go, single quotes are used to represent rune literals.
		A rune is Go's term for a single Unicode code point.


		'a' as a Rune:

		When we write 'a', we're creating a rune literal that represents the lowercase letter 'a'.
		In ASCII (and Unicode), 'a' has the numeric value 97.
		*/
		mask := uint64(1) << val
		if (checker & mask) > 0 {
			return true
		}
		checker |= mask
	}
	return false
}
