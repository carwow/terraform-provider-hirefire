package client

type UserResource struct {
	client *Client
}

type User struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	Notifications bool   `json:"notifications"`
}

type wrappedUser struct {
	User User `json:"user"`
}

func (r *UserResource) Get(id string) (*User, error) {
	wrapped := &wrappedUser{}
	err := r.client.getResource("users", id, wrapped)
	return &wrapped.User, err
}
