package main

import "fmt"

var memo map[int]int = make(map[int]int)

func main() {
	cnt := Climb(5)
	fmt.Println(cnt)
}

func Climb(numOfStairs int) int {

	if memo[numOfStairs] != 0 {
		return memo[numOfStairs]
	}

	if numOfStairs <= 0 || numOfStairs == 1 {
		return 1
	}

	memo[numOfStairs] = Climb(numOfStairs-1) + Climb(numOfStairs-2)

	return memo[numOfStairs]
}
