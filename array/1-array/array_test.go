package __array

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("Should return index of element", func(t *testing.T) {
		// arrange
		arr := []string{"cat", "dog", "bird", "fish"}
		x := "bird"
		want := 2

		// act
		got, err := Find(arr, x)

		// assert
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})
	t.Run("Should retun error not found element", func(t *testing.T) {
		// arrange
		arr := []string{"cat", "dog", "bird", "fish"}
		x := "elephant"
		want := 0
		wantErr := errors.New("not found")

		// act
		got, err := Find(arr, x)

		// assert
		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, err)
	})
}

func TestCounter(t *testing.T) {
	t.Run("should return element counter", func(t *testing.T) {
		// arrange
		arr := []string{"cat", "dog", "bird", "bird"}
		x := "bird"
		want := 2

		// act
		got, err := Counter(arr, x)

		// assert
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})
	t.Run("should return error element counter", func(t *testing.T) {
		// arrange
		arr := []string{"cat", "dog", "bird", "fish"}
		x := "elephant"
		want := 0
		wantErr := errors.New("not found")

		// act
		got, err := Counter(arr, x)

		// assert
		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, err)
	})
}
