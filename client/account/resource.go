package account

import "github.com/carwow/terraform-provider-hirefire/client/client"

type Resource struct {
	client client.Client
}

func NewResource(client client.Client) *Resource {
	return &Resource{client: client}
}

type Account struct {
	Id             string
	OrganizationId string `json:"organization_id"`
}

func (r *Resource) Get(id string) (*Account, error) {
	wrapped := &struct{ Account Account }{}
	err := r.client.GetResource("accounts", id, wrapped)
	return &wrapped.Account, err
}
