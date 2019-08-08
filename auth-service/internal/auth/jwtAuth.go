package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/franpog859/cleanaux-backend/auth-service/internal/kubernetes"
)

const (
	jwtTokenSecret = "lasdlkashdakjshd"
)

// CreateJWTToken function creates a JWT token using username and JTW secret
func CreateJWTToken(username string, kubernetesClient kubernetes.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	// TODO: use k8sClient to get jwtTokenSecret every time token is being created
	signedToken, err := token.SignedString([]byte(jwtTokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
