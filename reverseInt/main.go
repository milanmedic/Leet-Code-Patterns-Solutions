package main

import (
	"fmt"
	"strconv"
	"strings"
)

const INT32MAX = 2147483647
const INT32MIN = -2147483647

func main() {
	integer := reverse(-1534236469)
	fmt.Println(integer)
}

func reverse(integer int) int {

	if !CheckRange(integer) {
		return 0
	}

	var intStr string = ConvertToString(integer)

	if intStr[0] == '-' {
		var intStrWOSign string = intStr[1:]
		intStrWOSign = ReverseString(intStrWOSign)
		integer = ConvertToInt(intStrWOSign)
		integer = integer * -1
		if !CheckRange(integer) {
			return 0
		}
		return integer
	}

	intStr = ReverseString(intStr)
	integer = ConvertToInt(intStr)

	if !CheckRange(integer) {
		return 0
	}
	return integer
}

func CheckRange(integer int) bool {
	if integer >= INT32MIN && integer <= INT32MAX {
		return true
	}
	return false
}

func ConvertToString(integer int) string {
	return strconv.Itoa(integer)
}

func ReverseString(inputString string) string {
	intStrArr := strings.Split(inputString, "")
	for i, j := 0, len(inputString)-1; i < j; i, j = i+1, j-1 {
		intStrArr[i], intStrArr[j] = intStrArr[j], intStrArr[i]
	}
	intStr := strings.Join(intStrArr, "")
	return intStr
}

func ConvertToInt(inputString string) int {
	integer, err := strconv.Atoi(inputString)
	if err != nil {
		fmt.Println("There was an error")
	}
	return integer
}
