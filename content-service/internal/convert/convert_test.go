package convert

import (
	"testing"
	"time"

	"github.com/franpog859/cleanaux-backend/content-service/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvert_ContentFromItems(t *testing.T) {

	now := time.Now()

	t.Run("should return correct content from items", func(t *testing.T) {
		// given
		items := []model.Item{
			{
				ID:            1,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -9).Format(model.DateLayout),
				IntervalDays:  20,
			},
			{
				ID:            2,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -10).Format(model.DateLayout),
				IntervalDays:  20,
			},
			{
				ID:            3,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -14).Format(model.DateLayout),
				IntervalDays:  20,
			},
			{
				ID:            4,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -15).Format(model.DateLayout),
				IntervalDays:  20,
			},
			{
				ID:            5,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -18).Format(model.DateLayout),
				IntervalDays:  20,
			},
			{
				ID:            6,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -19).Format(model.DateLayout),
				IntervalDays:  20,
			},
			{
				ID:            7,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -3).Format(model.DateLayout),
				IntervalDays:  7,
			},
			{
				ID:            8,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -4).Format(model.DateLayout),
				IntervalDays:  7,
			},
			{
				ID:            9,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -5).Format(model.DateLayout),
				IntervalDays:  7,
			},
			{
				ID:            10,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -6).Format(model.DateLayout),
				IntervalDays:  7,
			},
			{
				ID:            11,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -7).Format(model.DateLayout),
				IntervalDays:  7,
			},
			{
				ID:            12,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, 0).Format(model.DateLayout),
				IntervalDays:  3,
			},
			{
				ID:            13,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -1).Format(model.DateLayout),
				IntervalDays:  3,
			},
			{
				ID:            14,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -2).Format(model.DateLayout),
				IntervalDays:  3,
			},
			{
				ID:            15,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -3).Format(model.DateLayout),
				IntervalDays:  3,
			},
			{
				ID:            16,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, 0).Format(model.DateLayout),
				IntervalDays:  1,
			},
			{
				ID:            17,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -1).Format(model.DateLayout),
				IntervalDays:  1,
			},
		}
		expectedContent := []model.ContentResponse{
			{ID: 1, Name: "name", Status: 0},
			{ID: 2, Name: "name", Status: 1},
			{ID: 3, Name: "name", Status: 1},
			{ID: 4, Name: "name", Status: 2},
			{ID: 5, Name: "name", Status: 2},
			{ID: 6, Name: "name", Status: 3},
			{ID: 7, Name: "name", Status: 0},
			{ID: 8, Name: "name", Status: 1},
			{ID: 9, Name: "name", Status: 1},
			{ID: 10, Name: "name", Status: 2},
			{ID: 11, Name: "name", Status: 3},
			{ID: 12, Name: "name", Status: 0},
			{ID: 13, Name: "name", Status: 0},
			{ID: 14, Name: "name", Status: 1},
			{ID: 15, Name: "name", Status: 3},
			{ID: 16, Name: "name", Status: 0},
			{ID: 17, Name: "name", Status: 3},
		}

		// when
		content, err := ContentFromItems(items)
		require.NoError(t, err)

		// then
		assert.Equal(t, expectedContent, content)
	})

	t.Run("should return error if data is intervalDays is incorrect", func(t *testing.T) {
		// given
		items := []model.Item{
			{
				ID:            1,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, -9).Format(model.DateLayout),
				IntervalDays:  0,
			},
		}

		// when
		_, err := ContentFromItems(items)

		// then
		assert.Error(t, err)
	})

	t.Run("should return error if time is incorrectly read", func(t *testing.T) {
		// given
		items := []model.Item{
			{
				ID:            1,
				Name:          "name",
				LastUsageDate: now.AddDate(0, 0, 20).Format(model.DateLayout),
				IntervalDays:  20,
			},
		}

		// when
		_, err := ContentFromItems(items)

		// then
		assert.Error(t, err)
	})
}

func TestConvert_UpdateItemInputFromContentRequest(t *testing.T) {

	now := time.Now()

	t.Run("should return correct update item", func(t *testing.T) {
		// given
		userRequestBody := model.ContentRequest{
			ID: 1,
		}
		expectedUpdateItem := model.UpdateItem{
			ID:            1,
			LastUsageDate: now.Format(model.DateLayout),
		}

		// when
		updateItem := UpdateItemFromContentRequest(userRequestBody)

		// then
		assert.Equal(t, expectedUpdateItem, updateItem)
	})
}
