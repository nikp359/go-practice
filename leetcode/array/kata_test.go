package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	input := []int{2, 3, 9, 10, 11}
	want := []int{11, 10, 9, 3, 2}

	got := reverse(input)
	assert.Equal(t, want, got)
}

// Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
// That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
// https://leetcode.com/problems/summary-ranges/description/
func TestSummary(t *testing.T) {

}
