package auth

import (
	"fmt"
	"time"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/cache"

	"github.com/dgrijalva/jwt-go"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
)

const (
	jwtTokenSecret       = "lasdlkashdakjshd"
	tokenExpirationHours = 5
)

// CreateJWTToken function creates a JWT token using username and JTW secret
func CreateJWTToken(username string, kubernetesClient kubernetes.Client) (string, error) {
	expirationTime := time.Now().Add(tokenExpirationHours * time.Hour)

	claims := &model.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// TODO: use k8sClient to get jwtTokenSecret every time token is being created
	signedToken, err := token.SignedString([]byte(jwtTokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// IsTokenValid function validates JWT token checking the signature and expiration time
func IsTokenValid(token string, kubernetesClient kubernetes.Client, tokenCache cache.Cache) (bool, error) {
	claims := &model.Claims{}

	// TODO: Use cache to save token
	// TODO: get jwtTokenSecret from secret if token is invalid
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtTokenSecret, nil
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
