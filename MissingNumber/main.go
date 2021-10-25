package main

import "fmt"

func main() {
	//[3,0,1]
	//[0,1]
	//[9,6,4,2,3,5,7,0,1]
	//[0]
	//[1]
	//[1,2]
	arr := []int{0, 1}
	fmt.Println(missingNumber(arr))
}

func missingNumber(nums []int) int {

	if len(nums) == 1 && nums[0] == 0 {
		return 1
	}

	if len(nums) == 1 && nums[0] == 1 {
		return 0
	}

	n := Max(nums)
	sum := n * (n + 1) / 2
	arrSum := Sum(nums)

	if sum-arrSum == 0 {

		if ContainsZero(nums) {
			return n + 1
		}
		return 0
	}

	return sum - arrSum
}

func ContainsZero(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return true
		}
	}
	return false
}

func Sum(nums []int) int {
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	return sum
}

func Max(nums []int) int {
	var currentMax int = -9999

	for i := 0; i < len(nums); i++ {
		if nums[i] > currentMax {
			currentMax = nums[i]
		}
	}

	return currentMax
}
