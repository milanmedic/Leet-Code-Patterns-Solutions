package main

import "fmt"

func main() {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums2 := []int{1}
	nums3 := []int{5, 4, -1, 7, 8}
	nums4 := []int{-2, 1}
	nums5 := []int{-2, -1, -3}
	nums6 := []int{-1, -2}
	nums7 := []int{1, 2}
	fmt.Println(MaxSubArr(nums1))
	fmt.Println(MaxSubArr(nums2))
	fmt.Println(MaxSubArr(nums3))
	fmt.Println(MaxSubArr(nums4))
	fmt.Println(MaxSubArr(nums5))
	fmt.Println(MaxSubArr(nums6))
	fmt.Println(MaxSubArr(nums7))
}

// Works for all but extremelly large cases
func MaxSubArr(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 0 {
		return 0
	}

	currSum, maxSum := -2147483648, -2147483648

	for i := 0; i < len(nums); i++ {
		currSum = nums[i]
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > currSum+nums[j] {
				i = j - 1
				break
			}
			currSum = currSum + nums[j]

			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}

	return maxSum
}
