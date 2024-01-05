package maxvowels

func maxVowels(s string, k int) int {
	vowels := countVowels(s[0:k])
	max := vowels

	if len(s) <= k+1 {
		return max
	}

	for i, j := k, k+1; j < len(s); i, j = i+1, j+1 {
		if isVowel(s[j]) {
			vowels++
		}

		if isVowel(s[i]) {
			vowels--
		}

		if vowels > max {
			max = vowels
		}
	}

	return max
}

func countVowels(s string) (v int) {
	for _, symbol := range s {
		if isVowel(byte(symbol)) {
			v++
		}
	}

	return v
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}
