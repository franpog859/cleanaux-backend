package database

import (
	"database/sql"
	"fmt"

	"github.com/franpog859/cleanaux-backend/content-service/internal/model"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseUsername = "root"
	databasePassword = "password"
	databaseBase     = "mysql-database-internal:3306"
	databaseName     = "content"
	driverName       = "mysql"
)

// Client interface
type Client interface {
	GetAllItems() ([]model.Item, error)
	UpdateItem(updateItem model.UpdateItem) error
	Close() error
}

type client struct {
	dbMySQL *sql.DB
}

// NewClient provides Client interface
func NewClient() (Client, error) {
	db, err := sql.Open(driverName, source())
	if err != nil {
		return nil, fmt.Errorf("failed to connect MySQL client: %v", err)
	}

	return &client{
		dbMySQL: db,
	}, nil
}

func source() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		databaseUsername,
		databasePassword,
		databaseBase,
		databaseName,
	)
}

func (c *client) GetAllItems() ([]model.Item, error) {
	selectDB, err := c.dbMySQL.Query("SELECT * FROM items")
	if err != nil {
		return []model.Item{}, err
	}
	defer selectDB.Close()

	items, err := getItemsFromQuery(selectDB)
	if err != nil {
		return []model.Item{}, err
	}

	return items, nil
}

func getItemsFromQuery(query *sql.Rows) ([]model.Item, error) {
	items := []model.Item{}
	var id, intervalDays int
	var name, lastUsageDate string

	for query.Next() {
		err := query.Scan(&id, &name, &lastUsageDate, &intervalDays)
		if err != nil {
			return []model.Item{}, err
		}

		itemInstance := model.Item{
			ID:            id,
			Name:          name,
			LastUsageDate: lastUsageDate,
			IntervalDays:  intervalDays,
		}
		items = append(items, itemInstance)
	}

	return items, nil
}

func (c *client) UpdateItem(updateItem model.UpdateItem) error {
	updateDB, err := c.dbMySQL.Prepare("UPDATE items SET lastUsageDate=? WHERE id=?")
	if err != nil {
		return err
	}
	defer updateDB.Close()

	_, err = updateDB.Exec(updateItem.LastUsageDate, updateItem.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Close() error {
	return c.dbMySQL.Close()
}
