package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert_GetContentFromItems(t *testing.T) {

	t.Run("should return correct content from items", func(t *testing.T) {
		// given
		items := []item{
			{1, "name1", 1, "2019-03-01", 7},
			{2, "name2", 2, "2019-03-02", 7},
		}
		expectedContent := []userContentResponse{
			{1, "name1", 2},
			{2, "name2", 2},
		}

		// when
		content := GetContentFromItems(items)

		// then
		assert.Equal(t, expectedContent, content)
	})
}
