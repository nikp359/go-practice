package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Input = `7 7
A  B  C
|  |  |
|--|  |
|  |--|
|  |--|
|  |  |
1  2  3`

func TestReadInput(t *testing.T) {
	want := &GhostLegs{
		Width:  7,
		Height: 7,
	}

	got := readInput(strings.NewReader(Input))
	assert.Equal(t, got, want)
}
