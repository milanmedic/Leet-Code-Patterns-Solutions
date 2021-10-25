package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	var minPrice int = 99999999
	var maxPrice int = 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			if prices[i]-minPrice > maxPrice {
				maxPrice = prices[i] - minPrice
			}
		}
	}
	return maxPrice
}
