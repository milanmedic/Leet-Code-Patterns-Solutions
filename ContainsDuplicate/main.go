package main

import "fmt"

var ocurrences map[int]int = make(map[int]int)

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if ocurrences[num] == 0 {
			ocurrences[num] = 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(arr))
}
