package main

var letterMapping map[string]string = map[string]string{
	"1":  "A",
	"2":  "B",
	"3":  "C",
	"4":  "D",
	"5":  "E",
	"6":  "F",
	"7":  "G",
	"8":  "H",
	"9":  "I",
	"10": "J",
	"11": "K",
	"12": "L",
	"13": "M",
	"14": "N",
	"15": "O",
	"16": "P",
	"17": "Q",
	"18": "R",
	"19": "S",
	"20": "T",
	"21": "U",
	"22": "V",
	"23": "W",
	"24": "X",
	"25": "Y",
	"26": "Z",
}

func main() {
	input := "11106"
	Decode(input)
}

func Decode(input string) int {
	// first pass the input string where the range is 2, meaning decode by grabbing two characters
	// second pass, grab one character at a time, if the character can't be decoded, take the previous character along with the one that couldn't be decoded, and then decode it.
	// third pass, opposite of second pass
	var cnt int = 0
	cnt = cnt + DecodeTwoCharRange(input)
	cnt = cnt + DecodeOneRangeDecrementRng(input)
	cnt = cnt + DecodeOneRangeIncrementRng(input)

	return cnt
}

func DecodeTwoCharRange(input string) int {
	var slice string
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) {
			slice = input[i : i+1]
		} else {
			slice = string(input[i])
		}
		if letterMapping[slice] == "" {
			return 0
		}
	}
	return 1
}

func DecodeOneRangeDecrementRng(input string) int {
	var slice string
	for i := 0; i < len(input); i++ {
		slice = string(input[i])
		if letterMapping[slice] == "" {
			if i-1 > 0 {
				slice = input[i-1 : i]
				if letterMapping[slice] == "" {
					return 0
				}
			}
		}
	}

	return 1
}

func DecodeOneRangeIncrementRng(input string) int {
	var slice string
	for i := 0; i < len(input); i++ {
		slice = string(input[i])
		if letterMapping[slice] == "" {
			if i+1 < len(input) {
				slice = input[i : i+1]
				if letterMapping[slice] == "" {
					return 0
				}
			}
		}
	}

	return 1
}
