package main

import (
	"fmt"
	"strings"
)

//	["flower","flow","flight"]
// "fl"

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// textArr := strings.Split(text, ",")
	arr := []string{"cir", "car"}
	fmt.Println(FindLongestCommonPrefix(arr))
}

func FindLongestCommonPrefix(allStrings []string) string {
	var prefix string = ""
	var smallestStrInd int
	prefix, smallestStrInd = FindSmallestStr(allStrings)

	allStrings = RemoveSmallestStr(smallestStrInd, allStrings)

	for i := 0; i < len(allStrings); i++ {
		for j := 0; j < len(prefix); j++ {
			if prefix[j] != allStrings[i][j] {
				prefix = RemoveCharacterFromPrefix(prefix)
				j--
			}
		}
	}

	return prefix
}

func FindSmallestStr(allStrings []string) (string, int) {
	var smallestStr string = allStrings[0]
	var smallestStrInd int = 0
	for ind, str := range allStrings {
		if len(smallestStr) > len(str) {
			smallestStr = str
			smallestStrInd = ind
		}
	}
	return smallestStr, smallestStrInd
}

func RemoveSmallestStr(index int, allStrings []string) []string {
	var newArr []string = []string{}
	newArr = append(newArr, allStrings[0:index]...)
	newArr = append(newArr, allStrings[index+1:]...)
	return newArr
}

func RemoveCharacterFromPrefix(prefix string) string {
	var splitPref []string = strings.Split(prefix, "")
	splitPref = splitPref[0 : len(splitPref)-1]
	return strings.Join(splitPref, "")
}
