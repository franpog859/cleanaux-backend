// +build integration

package main

import "testing"

var auth = authService{Base: "http://localhost:8001"}

func TestAuthService(t *testing.T) {
	t.Run("should not be able to login with wrong username/password", func(t) {

	})
}
