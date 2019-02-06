package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	searchText := "This is a test"
	dictonary := Dictonary{"test": searchText}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictonary.Search("test")
		want := searchText
		assert.Equal(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictonary.Search("unknown")

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrNotFound)
	})

}
