package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRanges(t *testing.T) {
	testCases := []struct {
		input []int
		want  []string
	}{
		{
			input: []int{0, 1, 2, 4, 5, 7},
			want:  []string{"0->2", "4->5", "7"},
		},
		{
			input: []int{0, 2, 3, 4, 6, 8, 9},
			want:  []string{"0", "2->4", "6", "8->9"},
		},
		{
			input: []int{0},
			want:  []string{"0"},
		},
		{
			input: []int{-1},
			want:  []string{"-1"},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case: %d", i), func(t *testing.T) {
			got := summaryRanges(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}

}
