package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
)

const (
	tokenExpirationHours = 5
	jwtBearer            = "Bearer"
)

// CreateJWTToken function creates a JWT token using username and JWT secret
func CreateJWTToken(username, JWTKey string) (string, error) {
	expirationTime := time.Now().Add(tokenExpirationHours * time.Hour)

	claims := &model.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ExtractTokenFromHeader provides token provided
// in the JWT Auth header and eventually an error
func ExtractTokenFromHeader(jwtAuthHeader string) (string, error) {
	splittedAuthHeader := strings.SplitN(jwtAuthHeader, " ", 2)
	if len(splittedAuthHeader) != 2 || splittedAuthHeader[0] != jwtBearer {
		return "", fmt.Errorf("invalid JWT Auth header structure: %s", jwtAuthHeader)
	}

	token := splittedAuthHeader[1]

	return token, nil
}

// IsTokenValid function validates JWT token checking
// the signature and expiration time
func IsTokenValid(token, JWTKey string) (bool, error) {
	claims := &model.Claims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, nil
		}
		return false, fmt.Errorf("failed to parse JWT token: %v", err)
	}
	if !parsedToken.Valid {
		return false, nil
	}

	return true, nil
}
