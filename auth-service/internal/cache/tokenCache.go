package cache

// Cache interface
type Cache interface {
}

type cache struct {
}

// New provides Cache interface
func New() Cache {
	return &cache{}
}
