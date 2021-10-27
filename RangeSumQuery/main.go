package main

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0

	if left > right {
		return sum
	}

	for i := left; i <= right; i++ {
		sum = sum + this.nums[i]
	}

	return sum
}

func main() {

}

/* MORE OPTIMAL SOLUTION
WHEN ADDING NUMS TO ARR, WE SAVE THE FIRST ONE AS IS, BUT WE ADD THE PREVIOUS ELEMENT WITH THE CURRENT ONE
THEN WE JUST SELECT THE RIGHT ELEMENT AND SUBSTRACT IT FROM THE LEFT ONE
func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums))
    sums[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        sums[i] = sums[i - 1] + nums[i]
    }

    return NumArray{sums}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left > 0 {
        return this.sums[right] - this.sums[left - 1]
    }
    return this.sums[right]
}

*/
