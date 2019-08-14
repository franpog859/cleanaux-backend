package cache

import (
	"github.com/franpog859/cleanaux-backend/auth-service/internal/model"
)

// Cache interface
type Cache interface {
	GetSecret() string
	SetSecret(string)
}

type cache struct {
}

// New provides Cache interface
func New() Cache {
	return &cache{}
}

func (c *cache) GetSecret() string {
	return model.JwtTokenSecret
}

func (c *cache) SetSecret(secret string) {
	return
}
