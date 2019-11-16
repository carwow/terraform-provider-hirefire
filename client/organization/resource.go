package organization

import "github.com/carwow/terraform-provider-hirefire/client/client"

type Resource struct {
	client client.Client
}

func NewResource(client client.Client) *Resource {
	return &Resource{client: client}
}

type Organization struct {
	Id       string
	Name     string
	TimeZone string `json:"time_zone"`
}

func (r *Resource) Get(id string) (*Organization, error) {
	wrapped := &struct{ Organization Organization }{}
	err := r.client.GetResource("organizations", id, wrapped)
	return &wrapped.Organization, err
}
