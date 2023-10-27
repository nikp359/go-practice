package candles

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := slices.Max(candies)

	response := make([]bool, 0, len(candies))

	for _, candle := range candies {
		greatest := false

		if candle+extraCandies >= max {
			greatest = true
		}

		response = append(response, greatest)
	}

	return response
}
