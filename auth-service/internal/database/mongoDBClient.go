package database

import "github.com/franpog859/cleanaux-backend/auth-service/internal/model"

var users = []model.User{
	{
		ID:       0,
		Username: "user0",
		Password: "pass0",
	},
	{
		ID:       1,
		Username: "user1",
		Password: "pass1",
	},
}

// Client interface
type Client interface {
	GetAllUsers() ([]model.User, error)
}

type client struct {
	mongoDBClient string
}

// NewClient provides Client interface
func NewClient() Client {
	return &client{}
}

func (c *client) GetAllUsers() ([]model.User, error) {
	return users, nil
}
