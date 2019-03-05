package main

import (
	"fmt"
	"testing"
)

func TestConvert_GetContentFromItems(t *testing.T) {

	t.Run("should no", func(t *testing.T) {
		items := []Item{
			Item{
				1, "name1", 1, "2019-03-01", 7,
			},
			Item{
				2, "name2", 2, "2019-03-02", 7,
			},
		}
		content := GetContentFromItems(items)
		fmt.Println(content)
	})
}
