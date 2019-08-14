package auth

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/franpog859/cleanaux-backend/auth-service/internal/database"
)

const (
	basicBearer = "Basic"
)

// ExtractCredentialsFromHeader provides username, password provided
// in the Basic Auth header and eventually an error
func ExtractCredentialsFromHeader(basicAuthHeader string) (string, string, error) {
	splittedAuthHeader := strings.SplitN(basicAuthHeader, " ", 2)
	if len(splittedAuthHeader) != 2 || splittedAuthHeader[0] != basicBearer {
		return "", "", fmt.Errorf("invalid Basic Auth header structure: %s", basicAuthHeader)
	}

	authPayload, _ := base64.StdEncoding.DecodeString(splittedAuthHeader[1])
	basicCredentials := strings.SplitN(string(authPayload), ":", 2)
	if len(basicCredentials) != 2 {
		return "", "", fmt.Errorf("invalid Basic Auth credentials structure: %s", authPayload)
	}

	username, password := basicCredentials[0], basicCredentials[1]

	return username, password, nil
}

// AreCredentialsValid validates user credentials checking users from the database
func AreCredentialsValid(username, password string, dbClient database.Client) (bool, error) {
	if err := validateCredentials(username, password); err != nil {
		return false, err
	}

	users, err := dbClient.GetAuthorizedUsers(username, password)
	if err != nil {
		return false, fmt.Errorf("failed to get all users from database: %v", err)
	}

	if len(users) < 1 {
		return false, nil
	}

	return true, nil
}

func validateCredentials(username, password string) error {
	regexPattern := regexp.MustCompile("^([a-zA-Z0-9]+)$")

	if !regexPattern.MatchString(username) {
		return fmt.Errorf("invalid username pattern: %s", username)
	}
	if !regexPattern.MatchString(password) {
		return fmt.Errorf("invalid password pattern: %s", password)
	}

	return nil
}
