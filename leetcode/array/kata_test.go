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

func TestKata(t *testing.T) {
	Kata()
}
