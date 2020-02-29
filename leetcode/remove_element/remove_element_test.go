package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElem(t *testing.T) {
	//cases := struct

	input := []int{1, 4, 1, 1, 1, 1, 1, 4, 5, 1, 97, 1, 1, 1, 5, 1}
	l := removeElement(input, 1)
	t.Log(input)
	assert.Equal(t, 5, l)
}
