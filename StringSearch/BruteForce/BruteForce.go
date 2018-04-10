package BruteForce

func brute(pattern string, textString string) int {
	patLen := len(pattern)
	textLen := len(textString)

	for i := 0; i <= textLen-patLen; i++ {
		j := 0
		for j < patLen && pattern[j] == textString[i+j] {
			j++
		}

		if j == patLen {
			return i
		}
	}

	return -1
}
