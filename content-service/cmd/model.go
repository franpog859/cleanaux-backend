package main

const dateLayout = "2006-01-02"

type item struct {
	ID            int
	Name          string
	LastUserID    int
	LastUsageDate string
	IntervalDays  int
}

type updateItem struct {
	ID            int
	LastUsageDate string
}

type userContentResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type userContentRequest struct {
	ID int `json:"id" binding:"required"`
}
