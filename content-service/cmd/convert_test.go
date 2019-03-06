package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConvert_GetContentFromItems(t *testing.T) {

	now := time.Now()
	layout := "2006-01-02"

	t.Run("should return correct content from items", func(t *testing.T) {
		// given
		items := []item{
			{
				1,
				"name",
				1,
				now.AddDate(0, 0, -9).Format(layout),
				20,
			},
			{
				2,
				"name",
				1,
				now.AddDate(0, 0, -10).Format(layout),
				20,
			},
			{
				3,
				"name",
				1,
				now.AddDate(0, 0, -14).Format(layout),
				20,
			},
			{
				4,
				"name",
				1,
				now.AddDate(0, 0, -15).Format(layout),
				20,
			},
			{
				5,
				"name",
				1,
				now.AddDate(0, 0, -18).Format(layout),
				20,
			},
			{
				6,
				"name",
				1,
				now.AddDate(0, 0, -19).Format(layout),
				20,
			},
			{
				7,
				"name",
				1,
				now.AddDate(0, 0, -3).Format(layout),
				7,
			},
			{
				8,
				"name",
				1,
				now.AddDate(0, 0, -4).Format(layout),
				7,
			},
			{
				9,
				"name",
				1,
				now.AddDate(0, 0, -5).Format(layout),
				7,
			},
			{
				10,
				"name",
				1,
				now.AddDate(0, 0, -6).Format(layout),
				7,
			},
			{
				11,
				"name",
				1,
				now.AddDate(0, 0, -7).Format(layout),
				7,
			},
			{
				12,
				"name",
				1,
				now.AddDate(0, 0, 0).Format(layout),
				3,
			},
			{
				13,
				"name",
				1,
				now.AddDate(0, 0, -1).Format(layout),
				3,
			},
			{
				14,
				"name",
				1,
				now.AddDate(0, 0, -2).Format(layout),
				3,
			},
			{
				15,
				"name",
				1,
				now.AddDate(0, 0, -3).Format(layout),
				3,
			},
			{
				16,
				"name",
				1,
				now.AddDate(0, 0, 0).Format(layout),
				1,
			},
			{
				17,
				"name",
				1,
				now.AddDate(0, 0, -1).Format(layout),
				1,
			},
		}
		expectedContent := []userContentResponse{
			{1, "name", 0},
			{2, "name", 1},
			{3, "name", 1},
			{4, "name", 2},
			{5, "name", 2},
			{6, "name", 3},
			{7, "name", 0},
			{8, "name", 1},
			{9, "name", 1},
			{10, "name", 2},
			{11, "name", 3},
			{12, "name", 0},
			{13, "name", 0},
			{14, "name", 1},
			{15, "name", 3},
			{16, "name", 0},
			{17, "name", 3},
		}

		// when
		content, err := GetContentFromItems(items)
		require.NoError(t, err)

		// then
		assert.Equal(t, expectedContent, content)
	})

	t.Run("should return error if data is intervalDays is incorrect", func(t *testing.T) {
		// given
		items := []item{
			{
				1,
				"name",
				1,
				now.AddDate(0, 0, -9).Format(layout),
				0,
			},
		}

		// when
		_, err := GetContentFromItems(items)

		// then
		require.Error(t, err)
	})

	t.Run("should return error if time is incorrectly read", func(t *testing.T) {
		// given
		items := []item{
			{
				1,
				"name",
				1,
				now.AddDate(0, 0, 20).Format(layout),
				20,
			},
		}

		// when
		_, err := GetContentFromItems(items)

		// then
		require.Error(t, err)
	})
}
