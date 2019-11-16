package client

type Client interface {
	GetResource(path string, id string, wrapped interface{}) error
}
