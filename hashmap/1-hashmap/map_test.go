package __hashmap

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	t.Run("Should return value of key", func(t *testing.T) {
		// arrange
		dictionary := map[string]string{"test": "this is just a test"}
		want := "this is just a test"
		// act
		got, err := Search(dictionary, "test")
		// assert
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})
	t.Run("Should return an error if key is not found", func(t *testing.T) {
		// arrange
		dictionary := map[string]string{"test": "this is just a test"}
		want := ""
		wantErr := errors.New("value not found")
		// act
		got, err := Search(dictionary, "testing")
		// assert
		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, err)
	})
}

func TestSearchMethod(t *testing.T) {
	t.Run("should return value of key", func(t *testing.T) {
		// arrange
		dictionary := dictionary{"test": "this is just a test"}
		want := "this is just a test"
		// act
		got, err := dictionary.Search("test")
		// assert
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})
	t.Run("should return error not found value of key", func(t *testing.T) {
		// arrange
		dictionary := dictionary{"test": "this is just a test"}
		want := ""
		wantErr := errors.New("value not found")
		// act
		got, err := dictionary.Search("testing")
		// assert
		assert.Equal(t, want, got)
		assert.Equal(t, wantErr, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should add key and value to dictionary", func(t *testing.T) {
		// arrange
		d := dictionary{}
		key := "test"
		value := "this is just a test"
		want := dictionary{"test": "this is just a test"}
		// act
		d.Add(key, value)
		// assert
		assert.Equal(t, want, d)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update value of key", func(t *testing.T) {
		// arrange
		d := dictionary{"test": "this is just a test"}
		key := "test"
		value := "this is just a test updated"
		want := dictionary{"test": "this is just a test updated"}
		// act
		err := d.Update(key, value)
		// assert
		assert.Equal(t, want, d)
		assert.Nil(t, err)
	})
	t.Run("shoud return error update key not exist", func(t *testing.T) {
		// arrange
		d := dictionary{"test": "this is just a test"}
		key := "testing"
		value := "this is just a test updated"
		want := dictionary{"test": "this is just a test"}
		wantErr := errors.New("key not found")
		// act
		err := d.Update(key, value)
		// assert
		assert.Equal(t, want, d)
		assert.Equal(t, wantErr, err)
	})
}
