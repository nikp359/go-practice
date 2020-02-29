package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		roman   string
		integer int
	}{
		{"", 0},
		{"III", 3},
	}

	for i, c := range cases {
		got := romanToInt(c.roman)
		assert.Equal(t, c.integer, got, fmt.Sprintf("Case %v, Roman %v", i, c.roman))
	}

}
