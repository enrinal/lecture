package go_testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	// arrange
	rectangle := Rectangle{10.0, 10.0}
	// act
	got := rectangle.Perimeter()
	want := 40.0

	// assert
	assert.Equal(t, want, got)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		areaRectangleTest := []struct {
			name  string
			shape Rectangle
			want  float64
		}{
			{
				name:  "Rectangle",
				shape: Rectangle{width: 12, height: 6},
				want:  72.0,
			},
			{
				name:  "Rectangle",
				shape: Rectangle{width: 5, height: 5},
				want:  25.0,
			},
		}

		for _, tt := range areaRectangleTest {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.shape.Area()
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("circles", func(t *testing.T) {
		// arrange
		circle := Circle{10.0}
		// act
		got := circle.Area()
		want := 314.1592653589793

		// assert
		assert.Equal(t, want, got, "they should be equal")
	})
}
