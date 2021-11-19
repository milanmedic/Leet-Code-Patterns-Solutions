package main

func main() {
	backspaceCompare("a#c", "b")
}

func backspaceCompare(s string, t string) bool {
	// push characters from first into stack
	// pop from stack when #
	// return stack, compare with other element for element
	sseq := CreateFinalSequence(s)
	tseq := CreateFinalSequence(t)

	if len(sseq) != len(tseq) {
		return false
	}

	for i := 0; i < len(sseq); i++ {
		if sseq[i] != tseq[i] {
			return false
		}
	}

	return true
}

func CreateFinalSequence(inputString string) []string {
	var sequence []string = []string{}
	for _, character := range inputString {
		if character != '#' {
			sequence = append(sequence, string(character))
		} else {
			if len(sequence) > 0 {
				sequence = sequence[:len(sequence)-1]
			}
		}
	}

	return sequence
}
