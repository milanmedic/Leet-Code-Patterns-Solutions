package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBits(2))
}

/*
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
Example 1:

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10


*/

var numOfOnesMap map[string]int = map[string]int{}

func countBits(n int) []int {
	//
	arr := []int{}
	for i := 0; i <= n; i++ {
		numOfOnes := 0
		binRep := strconv.FormatInt(int64(i), 2)

		_, found := numOfOnesMap[binRep]

		if !found {
			for j := 0; j < len(binRep); j++ {
				if binRep[j] == '1' {
					numOfOnes++
				}
			}
			numOfOnesMap[binRep] = numOfOnes
		} else {
			numOfOnes = numOfOnesMap[binRep]
		}

		arr = append(arr, numOfOnes)
	}

	return arr
}
