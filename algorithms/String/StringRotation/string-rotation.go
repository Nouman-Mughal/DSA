package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CheckStringRotation("waterbottle", "erbottlewat"))
	fmt.Println(CheckStringRotation("waterbottle", "rbottlewat"))

}

func CheckStringRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	concatenated := s1 + s1

	return strings.Contains(concatenated, s2)
}
