package go_testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaOne(t *testing.T) {
	t.Run("testing circle area", func(t *testing.T) {
		// arrange
		c := CircleOne{10}
		want := float64(314.1592653589793)
		// act
		got := c.Area()
		// assert
		assert.Equal(t, want, got)
	})
	t.Run("testing rectangle area", func(t *testing.T) {
		// arrange
		r := RectangleOne{10, 20}
		want := float64(200)
		// act
		got := r.Area()
		// assert
		assert.Equal(t, want, got)
	})
}

func TestAreaOneSlice(t *testing.T) {
	// arrange
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Circle",
			shape: &CircleOne{10},
			want:  float64(314.1592653589793),
		},
		{
			name:  "Rectangle",
			shape: &RectangleOne{10, 20},
			want:  float64(200),
		},
		{
			name:  "Circle#1",
			shape: &CircleOne{5},
			want:  float64(78.53981633974483),
		},
		{
			name:  "Rectangle#1",
			shape: &RectangleOne{5, 10},
			want:  float64(50),
		},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			// act
			got := tt.shape.Area()
			// assert
			assert.Equal(t, tt.want, got)
		})
	}
}
