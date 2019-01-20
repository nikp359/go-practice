package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectagle := Rectagle{10.0, 10.0}
	got := Perimeter(rectagle)
	want := 40.0

	assert.Equal(t, got, want, "they should be equal")
}

func TestArea(t *testing.T) {

	rectagle := Rectagle{10.0, 10.0}
	got := Area(rectagle)
	want := 100.0
	assert.Equal(t, got, want, "they should be equal")
}
