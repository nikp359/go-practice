package subsequence

func isSubsequence(s string, t string) bool {
	var i, j int

	for ; i < len(s); i++ {
		for {
			if j == len(t) {
				return false
			}

			if s[i] == t[j] {
				j++
				break
			}

			j++
		}
	}

	return true
}
