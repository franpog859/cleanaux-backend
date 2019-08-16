package model

const (
	// DateLayout is a layout used for dates.
	DateLayout = "2006-01-02"
)

// Item is a model of entity from mysql database
type Item struct {
	ID            int
	Name          string
	LastUsageDate string
	IntervalDays  int
}

// UpdateItem is used for updating item in mysql database
type UpdateItem struct {
	ID            int
	LastUsageDate string
}

// ContentResponse is a response provided by the /content GET endpoint
type ContentResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// ContentRequest is a request expected in the /content PUT endpoint
type ContentRequest struct {
	ID int `json:"id" binding:"required"`
}
