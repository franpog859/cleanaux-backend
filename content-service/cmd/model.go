package main

type item struct {
	ID            int
	Name          string
	LastUserID    int
	LastUsageDate string
	IntervalDays  int
}

type userContentResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
