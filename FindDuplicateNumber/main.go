package main

import (
	"fmt"
	"math"
)

func findDuplicate(nums []int) int {
	found := -1
	// use set or hashMap
	// negative marking solution

	for i := 0; i < len(nums); i++ {
		elem := int(math.Abs(float64(nums[i])))
		if nums[elem] < 0 {
			return elem
		}
		nums[elem] = 0 - nums[elem]

	}
	return found
}

func main() {
	arr := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(arr))
}
