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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectagle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assert.Equal(t, got, tt.want, "they should be equal")
	}
}
