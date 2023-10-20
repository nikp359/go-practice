package prefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	firstStr := strs[0]
	prefix := ""

	for _, char := range firstStr {
		checkPrfx := prefix + string(char)

		for _, str := range strs[1:] {
			if !strings.HasPrefix(str, checkPrfx) {
				return prefix
			}
		}

		prefix = checkPrfx
	}

	return prefix
}
