package kubernetes

import (
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
)

// Client interface
type Client interface {
	GetSecret() string
}

type client struct {
}

// NewClient provides Client interface
func NewClient() Client {
	return &client{}
}

func (c *client) GetSecret() string {
	return model.JwtTokenSecret
}
