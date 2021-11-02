package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	return BinarySearch(nums, left, right, target)
}

func BinarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if target == nums[mid] {
		return mid
	} else if target > nums[mid] {
		return BinarySearch(nums, mid+1, right, target)
	} else {
		return BinarySearch(nums, left, mid-1, target)
	}
}

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(arr, 2))
	fmt.Println(search(arr, 9))
	fmt.Println(search(arr, 0))
}
