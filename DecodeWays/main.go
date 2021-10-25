package main

import "fmt"

var decodedMsgs map[string]int = make(map[string]int)

func main() {
	input := "11126"
	fmt.Println(Decode(input))
}

func Decode(input string) int {
	if len(input) == 0 {
		return 1
	}

	if string(input[0]) == "0" {
		return 0
	}

	// if value, ok := decodedMsgs[input]; ok {
	// 	return value
	// }

	cnt := Decode(input[1:])
	if len(input) >= 2 {
		if input[0:2] <= "26" {
			cnt += Decode(input[2:])
		}
	}
	//decodedMsgs[input] = cnt

	return cnt
}
