package main

func main() {
	println(checkPermutation("hello", "olleh"))
}

func checkPermutation(s1 string, s2 string) bool {
	sum1 := 0
	sum2 := 0
	for _, char := range s1 {
		sum1 += int(char)
	}
	for _, char := range s2 {
		sum2 += int(char)
	}
	if sum1 == sum2 {
		return true
	}
	return false
}
