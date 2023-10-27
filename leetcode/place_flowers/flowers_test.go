package flowers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	testCases := []struct {
		flowerbed []int
		seedling  int
		want      bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1, 0, 0, 0, 1},
			seedling:  2,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 0, 0, 1},
			seedling:  2,
			want:      true,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 1},
			seedling:  1,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1, 0, 0},
			seedling:  2,
			want:      true,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 0},
			seedling:  1,
			want:      true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := canPlaceFlowers(tc.flowerbed, tc.seedling)
			assert.Equal(t, tc.want, got)
		})
	}
}
