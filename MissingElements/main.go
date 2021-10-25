package main

func main() {
	nums := []int{2, 2}
	findDisappearedNumbers(nums)
}

func findDisappearedNumbers(nums []int) []int {
	var allNums map[int]bool = make(map[int]bool)
	var output []int = []int{}

	for _, i := range nums {
		allNums[i] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !allNums[i] {
			output = append(output, i)
		}
	}

	return output
}
