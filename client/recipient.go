package client

type RecipientResource struct {
	client *Client
}

type Recipient struct {
	Id            string `json:"id"`
	ApplicationId string `json:"application_id"`
	Email         string `json:"email"`
}

type wrappedRecipient struct {
	Recipient Recipient `json:"recipient"`
}

func (r *RecipientResource) Get(id string) (*Recipient, error) {
	wrapped := &wrappedRecipient{}
	err := r.client.getResource("recipients", id, wrapped)
	return &wrapped.Recipient, err
}

func (r *RecipientResource) Create(create Recipient) (*Recipient, error) {
	wrapped := &wrappedRecipient{Recipient: create}
	err := r.client.createResource("recipients", wrapped)
	return &wrapped.Recipient, err
}

func (r *RecipientResource) Update(update Recipient) (*Recipient, error) {
	wrapped := &wrappedRecipient{Recipient: update}
	err := r.client.updateResource("recipients", update.Id, wrapped)
	return &wrapped.Recipient, err
}

func (r *RecipientResource) Delete(id string) error {
	err := r.client.deleteResource("recipients", id)
	return err
}
