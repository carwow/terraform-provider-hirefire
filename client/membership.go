package client

type MembershipResource struct {
	client *Client
}

type Membership struct {
	Id             string `json:"id"`
	OrganizationId string `json:"organization_id"`
	UserId         string `json:"user_id"`
	Owner          bool   `json:"owner"`
}

type wrappedMembership struct {
	Membership Membership `json:"membership"`
}

func (r *MembershipResource) Get(id string) (*Membership, error) {
	wrapped := &wrappedMembership{}
	err := r.client.getResource("memberships", id, wrapped)
	return &wrapped.Membership, err
}

func (r *MembershipResource) Create(create Membership) (*Membership, error) {
	wrapped := &wrappedMembership{Membership: create}
	err := r.client.createResource("memberships", wrapped)
	return &wrapped.Membership, err
}

func (r *MembershipResource) Update(update Membership) (*Membership, error) {
	wrapped := &wrappedMembership{Membership: update}
	err := r.client.updateResource("memberships", update.Id, wrapped)
	return &wrapped.Membership, err
}

func (r *MembershipResource) Delete(id string) error {
	err := r.client.deleteResource("memberships", id)
	return err
}
