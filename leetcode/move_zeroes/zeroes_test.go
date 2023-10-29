package move_zeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{0, 1, 0, 3, 12},
			want:  []int{1, 3, 12, 0, 0},
		},
		{
			input: []int{0},
			want:  []int{0},
		},
		{
			input: []int{0, 1},
			want:  []int{1, 0},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			moveZeroes(tc.input)
			assert.Equal(t, tc.want, tc.input)
		})
	}
}
