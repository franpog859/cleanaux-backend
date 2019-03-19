package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseUsername     = "root"
	databasePassword     = "password"
	databaseBase         = "mysql-database-internal:3306"
	databaseDatabaseName = "content"
)

// DatabaseService interface.
type DatabaseService interface {
	GetAllItems() ([]item, error)
	UpdateItem(updateItem updateItem) error
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

func (database *databaseService) GetAllItems() ([]item, error) {
	db, err := sql.Open("mysql", database.source())
	if err != nil {
		return []item{}, err
	}
	defer db.Close()

	selectDB, err := db.Query("SELECT * FROM items")
	if err != nil {
		return []item{}, err
	}

	items, err := getItemsFromQuery(selectDB)
	if err != nil {
		return []item{}, err
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

func getItemsFromQuery(query *sql.Rows) ([]item, error) {
	items := []item{}
	var id, intervalDays, lastUserID int
	var name, lastUsageDate string

	for query.Next() {
		err := query.Scan(&id, &name, &lastUserID, &lastUsageDate, &intervalDays)
		if err != nil {
			return []item{}, err
		}

		itemInstance := item{
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

func (database *databaseService) UpdateItem(updateItem updateItem) error {
	db, err := sql.Open("mysql", database.source())
	if err != nil {
		return err
	}
	defer db.Close()

	updateDB, err := db.Prepare("UPDATE items SET lastUsageDate=? WHERE id=?")
	if err != nil {
		return err
	}

	_, err = updateDB.Exec(updateItem.LastUsageDate, updateItem.ID)
	if err != nil {
		return err
	}

	return nil
}
