package countOnes

/*
Suppose i = 5 (binary: 101).

i >> 1 shifts 101 to 10 (binary for 2), which has 1 one in it.
The LSB of 5 (binary 101) is 1, so we add 1 to the count of 1s in 2.
Result: countBits(5) = countBits(2) + 1 = 2.
*/

func countOnes(n int) []int {
	result := make([]int, n+1)

	for i:=1; i<=n;i++ {
		result[i] = result[i >> 1] + (i & 1)
	}
	return result
}