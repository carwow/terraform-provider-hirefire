package client

type ApplicationResource struct {
	client *Client
}

type Application struct {
	Id                         string  `json:"id"`
	AccountId                  string  `json:"account_id"`
	Name                       string  `json:"name"`
	Token                      string  `json:"token"`
	CustomDomain               *string `json:"custom_domain"`
	LogplexDrainToken          *string `json:"logplex_drain_token"`
	Ssl                        bool    `json:"ssl"`
	RestartCrashedDynos        bool    `json:"restart_crashed_dynos"`
	NewIssueNotifications      bool    `json:"new_issue_notifications"`
	ResolvedIssueNotifications bool    `json:"resolved_issue_notifications"`
	CheckupFrequency           int     `json:"checkup_frequency"`
}

type wrappedApplication struct {
	Application Application `json:"application"`
}

func (r *ApplicationResource) Get(id string) (*Application, error) {
	wrapped := &wrappedApplication{}
	err := r.client.getResource("applications", id, wrapped)
	return &wrapped.Application, err
}

func (r *ApplicationResource) Create(create Application) (*Application, error) {
	wrapped := &wrappedApplication{Application: create}
	err := r.client.createResource("applications", wrapped)
	return &wrapped.Application, err
}

func (r *ApplicationResource) Update(update Application) (*Application, error) {
	wrapped := &wrappedApplication{Application: update}
	err := r.client.updateResource("applications", update.Id, wrapped)
	return &wrapped.Application, err
}

func (r *ApplicationResource) Delete(id string) error {
	err := r.client.deleteResource("applications", id)
	return err
}
