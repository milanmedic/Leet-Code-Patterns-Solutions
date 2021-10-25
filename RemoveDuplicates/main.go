package main

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(arr)
}

// DOESN'T WORK
func removeDuplicates(nums []int) int {
	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return len(nums[:i+1])
}
