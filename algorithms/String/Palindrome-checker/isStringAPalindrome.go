// using two Pointer approach.
package palindromechecker

func isPalindrome(s string) bool {
    

    for left, right := 0, len(s)-1;left < right; {
        // Skip non-alphanumeric characters from the left
        if !isAlphanumeric(s[left]) {
            left++
            continue
        }

        // Skip non-alphanumeric characters from the right
        if !isAlphanumeric(s[right]) {
            right--
            continue
        }

        // Compare characters, ignoring case
        if toLowerCase(s[left]) != toLowerCase(s[right]) {
            return false
        }

        left++
        right--
    }

    return true
}

func isAlphanumeric(char byte) bool {
    return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toLowerCase(char byte) byte {
    if char >= 'A' && char <= 'Z' {
        return char + 32
    }
    return char
}
