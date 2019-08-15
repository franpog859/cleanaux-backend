package secret

import (
	"fmt"
	"io/ioutil"
)

const (
	secretPath = "/data/secret/jwtkey"
)

// Get returns the content of mounted secret file and an error if occurred
func Get() (string, error) {
	secret, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return "", fmt.Errorf("failed to read secret file: %v", err)
	}
	if len(secret) < 1 {
		return "", fmt.Errorf("secret is empty")
	}

	return string(secret), nil
}
