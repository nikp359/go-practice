package altitude

func largestAltitude(gain []int) int {
	var la, current int

	for _, a := range gain {
		current += a

		if current > la {
			la = current
		}
	}

	return la
}
