package client

type OrganizationResource struct {
	client *Client
}

type Organization struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	TimeZone string `json:"time_zone"`
}

type wrappedOrganization struct {
	Organization Organization `json:"organization"`
}

func (r *OrganizationResource) Get(id string) (*Organization, error) {
	wrapped := &wrappedOrganization{}
	err := r.client.getResource("organizations", id, wrapped)
	return &wrapped.Organization, err
}

func (r *OrganizationResource) Create(input Organization) (*Organization, error) {
	wrapped := &wrappedOrganization{Organization: input}
	err := r.client.createResource("organizations", wrapped)
	return &wrapped.Organization, err
}

func (r *OrganizationResource) Delete(id string) error {
	err := r.client.deleteResource("organizations", id)
	return err
}
