package auth

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	basicBearer = "Basic"
)

// ExtractCredentialsFromHeader provides username, password provided
// in the Basic Auth header and eventually an error
func ExtractCredentialsFromHeader(basicAuthHeader string) (string, string, error) {
	splittedAuthHeader := strings.SplitN(basicAuthHeader, " ", 2)

	if len(splittedAuthHeader) != 2 || splittedAuthHeader[0] != basicBearer {
		return "", "", fmt.Errorf("invalid Basic Auth header: %s", basicAuthHeader)
	}

	authPayload, _ := base64.StdEncoding.DecodeString(splittedAuthHeader[1])
	basicCredentials := strings.SplitN(string(authPayload), ":", 2)

	if len(basicCredentials) != 2 {
		return "", "", fmt.Errorf("invalid Basic Auth credentials: %s", authPayload)
	}

	username, password := basicCredentials[0], basicCredentials[1]

	return username, password, nil
}
