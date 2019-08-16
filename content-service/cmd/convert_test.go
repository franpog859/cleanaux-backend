package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvert_CreateContentFromItems(t *testing.T) {

	now := time.Now()

	t.Run("should return correct content from items", func(t *testing.T) {
		// given
		items := []item{
			{
				1,
				"name",
				now.AddDate(0, 0, -9).Format(dateLayout),
				20,
			},
			{
				2,
				"name",
				now.AddDate(0, 0, -10).Format(dateLayout),
				20,
			},
			{
				3,
				"name",
				now.AddDate(0, 0, -14).Format(dateLayout),
				20,
			},
			{
				4,
				"name",
				now.AddDate(0, 0, -15).Format(dateLayout),
				20,
			},
			{
				5,
				"name",
				now.AddDate(0, 0, -18).Format(dateLayout),
				20,
			},
			{
				6,
				"name",
				now.AddDate(0, 0, -19).Format(dateLayout),
				20,
			},
			{
				7,
				"name",
				now.AddDate(0, 0, -3).Format(dateLayout),
				7,
			},
			{
				8,
				"name",
				now.AddDate(0, 0, -4).Format(dateLayout),
				7,
			},
			{
				9,
				"name",
				now.AddDate(0, 0, -5).Format(dateLayout),
				7,
			},
			{
				10,
				"name",
				now.AddDate(0, 0, -6).Format(dateLayout),
				7,
			},
			{
				11,
				"name",
				now.AddDate(0, 0, -7).Format(dateLayout),
				7,
			},
			{
				12,
				"name",
				now.AddDate(0, 0, 0).Format(dateLayout),
				3,
			},
			{
				13,
				"name",
				now.AddDate(0, 0, -1).Format(dateLayout),
				3,
			},
			{
				14,
				"name",
				now.AddDate(0, 0, -2).Format(dateLayout),
				3,
			},
			{
				15,
				"name",
				now.AddDate(0, 0, -3).Format(dateLayout),
				3,
			},
			{
				16,
				"name",
				now.AddDate(0, 0, 0).Format(dateLayout),
				1,
			},
			{
				17,
				"name",
				now.AddDate(0, 0, -1).Format(dateLayout),
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
		content, err := CreateContentFromItems(items)
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
				now.AddDate(0, 0, -9).Format(dateLayout),
				0,
			},
		}

		// when
		_, err := CreateContentFromItems(items)

		// then
		assert.Error(t, err)
	})

	t.Run("should return error if time is incorrectly read", func(t *testing.T) {
		// given
		items := []item{
			{
				1,
				"name",
				now.AddDate(0, 0, 20).Format(dateLayout),
				20,
			},
		}

		// when
		_, err := CreateContentFromItems(items)

		// then
		assert.Error(t, err)
	})
}

func TestConvert_CreateUpdateItemInput(t *testing.T) {

	now := time.Now()

	t.Run("should return correct update item", func(t *testing.T) {
		// given
		userRequestBody := userContentRequest{
			ID: 1,
		}
		expectedUpdateItem := updateItem{
			ID:            1,
			LastUsageDate: now.Format(dateLayout),
		}

		// when
		updateItem := CreateUpdateItemInput(userRequestBody)

		// then
		assert.Equal(t, expectedUpdateItem, updateItem)
	})
}
