package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{2, 4, 7, 11, 13, 15}, 26, []int{3, 5}},
	}

	for _, tc := range testCases {
		got := twoSum(tc.nums, tc.target)
		assert.Equal(t, tc.want, got)

		got = twoSumHash(tc.nums, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
