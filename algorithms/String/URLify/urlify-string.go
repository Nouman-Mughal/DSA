// write an algorithm to replace all the spaces in a string with %20.
package main

import "fmt"

func main() {
	fmt.Println(urlify("hello world this is harsh world"))
}

// // this algorithm did not work.
// func urlify(s string) string {
// 	arr := []rune(s) // this array has fixed length equal to lenght of string.
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == ' ' {
// 			arr[i] = '%'
// 		}

//		}
//		return string(arr)
//	}
func urlify(s string) string {
	spaceCount := 0
	for _, r := range s {
		if r == ' ' {
			spaceCount++
		}
	}

	result := make([]rune, len(s)+spaceCount*3)
	j := 0

	// for index, value := range collection {
	// loop body
	// }
	for _, r := range s {
		if r == ' ' {
			result[j] = '%'
			result[j+1] = '2'
			result[j+2] = '0'
			j += 3
		} else {
			result[j] = r
			j++
		}
	}
	return string(result)
}
