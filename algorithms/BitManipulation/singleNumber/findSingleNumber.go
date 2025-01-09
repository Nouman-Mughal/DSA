package main

import "fmt"

func singleNumber(nums ...int) int {
    var result int
	// we need to return single number that does not appear twice.
    for _, num := range nums {
        result ^= num
		fmt.Println(result)
    }
    return result
}
func main(){
	singleNumber(2,2,1)
}