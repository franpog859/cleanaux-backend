package model

import "github.com/dgrijalva/jwt-go"

// User is a model of entity from mongo database
type User struct {
	Username string
	Password string
}

// TokenResponse is a response provided by the /login endpoint
type TokenResponse struct {
	Token string `json:"token"`
}

// Claims struct used as StandardClaims for generating JWT token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const (
	// AuthHeaderKey is a header key used for authorization
	AuthHeaderKey = "Authorization"

	// JwtTokenSecret is a mocked secret. TODO: Delete it after implementing real one!
	JwtTokenSecret = "lasdlkashdakjshd"
)
