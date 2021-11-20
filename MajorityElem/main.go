func majorityElement(nums []int) int {
	var elems map[int]int = make(map[int]int)

	for _, element := range nums {
		if _, found := elems[element]; found {
			elems[element] = elems[element] + 1
		} else {
			elems[element] = 1
		}
	}

	currMax := 0
	currKey := 0
	for key, value := range elems {
		if value > currMax {
			currMax = value
			currKey = key
		}
	}

	return currKey
}

func majorityElement(nums []int) int {
	cnt := 0
	candidate := -99999999

	for _, element := range nums {
		if cnt == 0 {
			candidate = element
		}

		if element == candidate {
			cnt += 1
		} else {
			cnt -= 1
		}
	}

	return candidate
}