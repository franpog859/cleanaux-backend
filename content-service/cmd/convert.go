package main

import "time"

func GetContentFromItems(items []Item) []Content {
	content := []Content{}
	for _, item := range items {
		status := calculateStatus(item)
		c := Content{
			item.ID,
			item.Name,
			status,
		}
		content = append(content, c)
	}

	return content
}

func calculateStatus(item Item) int {
	lastUsageDate := parseDate(item.LastUsageDate)

	t := time.Now()
	currentDate := t.Format("2006-01-02")
	days := currentDate.Sub(lastUsageDate).Hours() / 24
	return 2
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func parseDate(date string) time.Time {
	layout := "2006-01-02"
	return time.Parse(layout, date)
}
