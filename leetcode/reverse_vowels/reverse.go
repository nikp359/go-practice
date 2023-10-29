package reverse

func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}

	response := []byte(s)
	j := len(s) - 1

	for i := 0; i < len(s); i++ {
		if i >= j {
			return string(response)
		}

		if isVowel(s[i]) {
			for ; j >= 1; j-- {
				if i >= j {
					return string(response)
				}

				if isVowel(s[j]) {
					response[i], response[j] = response[j], response[i]
					j--
					break
				}
			}
		}
	}

	return string(response)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}
