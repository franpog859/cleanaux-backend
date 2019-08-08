package kubernetes

// Client interface
type Client interface {
}

type client struct {
}

// NewClient provides Client interface
func NewClient() Client {
	return &client{}
}
