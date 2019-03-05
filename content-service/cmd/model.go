package main

type Item struct {
	ID            int
	Name          string
	LastUserID    int
	LastUsageDate string
	IntervalDays  int
}

type Content struct {
	ID     int
	Name   string
	Status int
}
