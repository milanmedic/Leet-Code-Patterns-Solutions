package main

import (
	"fmt"
	"math"
)

func peakIndexInMountainArray(arr []int) int {

	return BinarySearch(0, len(arr)-1, arr)
}

func BinarySearch(left int, right int, arr []int) int {
	middle := int(math.RoundToEven(float64((left + right) / 2)))

	if left >= 0 && right <= len(arr)-1 {
		if arr[middle-1] < arr[middle] && arr[middle] > arr[middle+1] {
			return middle
		} else if arr[middle-1] > arr[middle] {
			return BinarySearch(0, middle, arr)
		} else if arr[middle+1] > arr[middle] {
			return BinarySearch(middle, right, arr)
		}
	}

	return middle
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
	fmt.Println(peakIndexInMountainArray([]int{3, 4, 5, 1}))
	fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
	fmt.Println(peakIndexInMountainArray([]int{40, 48, 61, 75, 100, 99, 98, 39, 30, 10}))
	fmt.Println(peakIndexInMountainArray([]int{55, 59, 63, 99, 97, 94, 84, 81, 79, 66, 40, 38, 33, 23, 22, 21, 17, 9, 7}))
}
