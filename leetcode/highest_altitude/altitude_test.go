package altitude

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	testCases := []struct {
		gain []int
		want int
	}{
		{
			gain: []int{-5, 1, 5, 0, -7},
			want: 1,
		},
		{
			gain: []int{-4, -3, -2, -1, 4, 3, 2},
			want: 0,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case n: %d", i+1), func(t *testing.T) {
			got := largestAltitude(tc.gain)
			assert.Equal(t, tc.want, got)
		})
	}
}
