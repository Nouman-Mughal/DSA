package countSingles

// finding the number that appear only once where all other appear twice.
// using method of frequency map
func singleNumberThree(nums []int) []int {
    freq := make(map[int]int)

    // Count the frequency of each number
    for _, num := range nums {
        freq[num]++
    }

    // Collect numbers that appear only once
    result := []int{}
    for num, count := range freq {
        if count == 1 {
            result = append(result, num)
        }
    }

    return result
}

