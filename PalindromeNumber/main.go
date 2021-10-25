package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	var intStr string = strconv.Itoa(x)
	var middle int = len(intStr) / 2

	for i, j := 0, len(intStr)-1; i < middle && j >= middle; i, j = i+1, j-1 {
		if intStr[i] != intStr[j] {

			return false
		}
	}

	return true
}

// if num is 123, take 1000, detract 123, result should then be reversed and summed with 123 if result is 1000, then true
