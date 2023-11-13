package reversewords

import "strings"

func reverseWords(s string) string {
	var resp strings.Builder

	sl := strings.Split(strings.Trim(s, " "), " ")
	resp.WriteString(sl[len(sl)-1])

	if len(sl) == 1 {
		return resp.String()
	}

	for i := len(sl) - 2; i >= 0; i-- {
		if sl[i] == "" {
			continue
		}

		resp.WriteString(" ")
		resp.WriteString(sl[i])
	}

	return resp.String()
}
