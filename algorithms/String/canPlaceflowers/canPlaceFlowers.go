package main

import "fmt"

func main() {
	arr := []int{0, 0, 0, 0, 1}
	println(canPlaceFlowers(arr, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if (i == len(flowerbed)-1 || flowerbed[i+1] == 0) && (i == 0 || flowerbed[i-1] == 0) {
				flowerbed[i] = 1
				n--
			}
		}
		fmt.Println(flowerbed)
	}
	return n <= 0
}
