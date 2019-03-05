package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseUsername     = "root"
	databasePassword     = "password"
	databaseBase         = "mysql:3306"
	databaseDatabaseName = "content"
)

// DatabaseService interface.
type DatabaseService interface {
	GetAllItems() ([]Item, error)
}

type databaseService struct {
	Username     string
	Password     string
	Base         string
	DatabaseName string
}

// NewDatabaseService provides DatabaseService interface.
func NewDatabaseService() DatabaseService {
	return &databaseService{
		databaseUsername,
		databasePassword,
		databaseBase,
		databaseDatabaseName,
	}
}

func (database *databaseService) GetAllItems() ([]Item, error) {
	db, err := sql.Open("mysql", database.source())
	if err != nil {
		return []Item{}, err
	}
	defer db.Close()

	selectDB, err := db.Query("SELECT * FROM items")
	if err != nil {
		return []Item{}, err
	}

	items, err := getItemsFromQuery(selectDB)
	if err != nil {
		return []Item{}, err
	}

	return items, nil
}

func (database *databaseService) source() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		database.Username,
		database.Password,
		database.Base,
		database.DatabaseName,
	)
}

func getItemsFromQuery(query *sql.Rows) ([]Item, error) {
	items := []Item{}
	var id, intervalDays, lastUserID int
	var name, lastUsageDate string

	for query.Next() {
		err := query.Scan(&id, &name, &lastUserID, &lastUsageDate, &intervalDays)
		if err != nil {
			return []Item{}, err
		}

		itemInstance := Item{
			id,
			name,
			lastUserID,
			lastUsageDate,
			intervalDays,
		}
		items = append(items, itemInstance)
	}

	return items, nil
}
