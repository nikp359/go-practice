package candles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:         []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:         []bool{true, false, false, false, false},
		},
		{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			want:         []bool{true, false, true},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := kidsWithCandies(tc.candies, tc.extraCandies)
			assert.Equal(t, tc.want, got)
		})
	}
}
