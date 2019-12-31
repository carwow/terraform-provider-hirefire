package client

type AccountResource struct {
	client *Client
}

type Account struct {
	Id             string `json:"id"`
	OrganizationId string `json:"organization_id"`
}

type wrappedAccount struct {
	Account Account `json:"account"`
}

func (r *AccountResource) Get(id string) (*Account, error) {
	wrapped := &wrappedAccount{}
	err := r.client.getResource("accounts", id, wrapped)
	return &wrapped.Account, err
}

func (r *AccountResource) Create(create Account) (*Account, error) {
	wrapped := &wrappedAccount{Account: create}
	err := r.client.createResource("accounts", wrapped)
	return &wrapped.Account, err
}

func (r *AccountResource) Update(update Account) (*Account, error) {
	wrapped := &wrappedAccount{Account: update}
	err := r.client.updateResource("accounts", update.Id, wrapped)
	return &wrapped.Account, err
}

func (r *AccountResource) Delete(id string) error {
	err := r.client.deleteResource("accounts", id)
	return err
}
