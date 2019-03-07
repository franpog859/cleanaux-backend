package main

import (
	"errors"
	"fmt"
	"time"
)

func CreateContentFromItems(items []item) ([]userContentResponse, error) {
	content := []userContentResponse{}
	for _, item := range items {
		status, err := getStatus(item)
		if err != nil {
			return content, err
		}

		c := userContentResponse{
			ID:     item.ID,
			Name:   item.Name,
			Status: status,
		}
		content = append(content, c)
	}

	return content, nil
}

func getStatus(item item) (int, error) {
	pastDays, err := countDays(item.LastUsageDate)
	if err != nil {
		return 0, err
	}

	status, err := calculateStatus(item.IntervalDays, pastDays)
	if err != nil {
		return 0, err
	}

	return status, nil
}

func countDays(lastUsageDate string) (int, error) {
	lastUsage, err := time.Parse(DATE_LAYOUT, lastUsageDate)
	if err != nil {
		return 0, err
	}

	pastDays := int(time.Now().Sub(lastUsage).Hours() / 24)
	return pastDays, nil
}

func calculateStatus(intervalDays, pastDays int) (int, error) {
	if intervalDays < 1 || pastDays < 0 {
		return 0, errors.New(fmt.Sprintf(
			"error while calculating status from intervalDays: %d and pastDays: %d",
			intervalDays, pastDays,
		))
	}

	var border int
	if intervalDays < 40 {
		border = intervalDays / 20
	} else {
		border = 2
	}
	if border >= intervalDays-pastDays {
		return 3, nil
	}

	if intervalDays < 20 {
		border = intervalDays / 4
	} else {
		border = 5
	}
	if border >= intervalDays-pastDays {
		return 2, nil
	}

	if intervalDays < 20 {
		border = intervalDays / 2
	} else {
		border = 10
	}
	if border >= intervalDays-pastDays {
		return 1, nil
	}

	return 0, nil
}

func CreateUpdateItemInput(userContentRequestBody userContentRequest) updateItem {
	lastUsageDate := time.Now().Format(DATE_LAYOUT)

	return updateItem{
		ID:            userContentRequestBody.ID,
		LastUsageDate: lastUsageDate,
	}
}
