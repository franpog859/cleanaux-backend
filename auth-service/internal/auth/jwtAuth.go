package auth

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/cache"

	"github.com/dgrijalva/jwt-go"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
)

const (
	tokenExpirationHours = 5
	jwtBearer            = "Bearer"
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
	jwtTokenSecret := kubernetesClient.GetSecret()

	signedToken, err := token.SignedString([]byte(jwtTokenSecret))
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

// IsTokenValid function validates JWT token checking the signature and expiration time
func IsTokenValid(token string, tokenCache cache.Cache, kubernetesClient kubernetes.Client) (bool, error) {
	// TODO: Refactor this function after implementing interfaces
	jwtTokenSecret := tokenCache.GetSecret()

	valid, _ := parseToken(token, jwtTokenSecret)
	if !valid {
		log.Printf("Failed to parse JWT token with cached secret. Retrying with Kubernetes Secret...")

		jwtTokenSecret := kubernetesClient.GetSecret()
		tokenCache.SetSecret(jwtTokenSecret)

		valid, err := parseToken(token, jwtTokenSecret)
		if err != nil {
			log.Printf("Invalid JWT token: %v", err)
			return false, nil
		}
		if !valid {
			return false, nil
		}
	}

	return true, nil
}

func parseToken(token, secret string) (bool, error) {
	claims := &model.Claims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
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
