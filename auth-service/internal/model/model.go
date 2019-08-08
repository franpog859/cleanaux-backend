package model

// User is a model of entity from mongo database
type User struct {
	ID       int
	Username string
	Password string
}

// TokenResponse is a response provided by the /login endpoint
type TokenResponse struct {
	Token string `json:"token"`
}
