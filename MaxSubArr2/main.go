package main

import "fmt"

func main() {
	var nums = []int{-1, -2, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(MaxSubArr(nums))
}

func MaxSubArr(arr []int) int {
	return CalculateMaxSubArr(arr, -99999999)
}

func CalculateMaxSubArr(arr []int, maxSum int) int {

	if len(arr) <= 0 {
		return maxSum
	}

	currSum := arr[0]

	if maxSum < currSum {
		maxSum = currSum
	}

	for i := 1; i < len(arr); i++ {
		currSum = currSum + arr[i]
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return CalculateMaxSubArr(arr[1:], maxSum)
}
