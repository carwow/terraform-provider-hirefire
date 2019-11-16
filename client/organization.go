package client

type OrganizationResource struct {
	client *Client
}

type Organization struct {
	Id       string
	Name     string
	TimeZone string `json:"time_zone"`
}

type wrappedOrganization struct{ Organization Organization }

func (r *OrganizationResource) Get(id string) (*Organization, error) {
	wrapped := &wrappedOrganization{}
	err := r.client.getResource("organizations", id, wrapped)
	return &wrapped.Organization, err
}
