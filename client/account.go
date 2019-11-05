package client

type AccountResource struct {
	client *Client
}

type Account struct {
	Id             string
	OrganizationId string `json:"organization_id"`
}

type wrappedAccount struct{ Account Account }

func (r *AccountResource) Get(id string) (*Account, error) {
	wrapped := &wrappedAccount{}
	err := r.client.getResource("accounts", id, wrapped)
	return &wrapped.Account, err
}
