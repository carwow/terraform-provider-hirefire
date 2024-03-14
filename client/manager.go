package client

type ManagerResource struct {
	client *Client
}

type Manager struct {
	Id            string `json:"id"`
	ApplicationId string `json:"application_id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Enabled       bool   `json:"enabled"`
	Minimum       int    `json:"minimum"`
	Maximum       int    `json:"maximum"`

	Aggregation          *string `json:"aggregation"`
	Percentile           *int    `json:"percentile"`
	MinimumLatency       *int    `json:"minimum_latency"`
	MaximumLatency       *int    `json:"maximum_latency"`
	MinimumQueueTime     *int    `json:"minimum_queue_time"`
	MaximumQueueTime     *int    `json:"maximum_queue_time"`
	MinimumResponseTime  *int    `json:"minimum_response_time"`
	MaximumResponseTime  *int    `json:"maximum_response_time"`
	MinimumConnectTime   *int    `json:"minimum_connect_time"`
	MaximumConnectTime   *int    `json:"maximum_connect_time"`
	MinimumLoad          *int    `json:"minimum_load"`
	MaximumLoad          *int    `json:"maximum_load"`
	MinimumApdex         *int    `json:"minimum_apdex"`
	MaximumApdex         *int    `json:"maximum_apdex"`
	LastMinutes          *int    `json:"last_minutes"`
	Ratio                *int    `json:"ratio"`
	Decrementable        *bool   `json:"decrementable"`
	Url                  *string `json:"url"`
	UpscaleQuantity      *int    `json:"upscale_quantity"`
	DownscaleQuantity    *int    `json:"downscale_quantity"`
	UpscaleSensitivity   *int    `json:"upscale_sensitivity"`
	DownscaleSensitivity *int    `json:"downscale_sensitivity"`
	UpscaleTimeout       int     `json:"upscale_timeout"`
	DownscaleTimeout     int     `json:"downscale_timeout"`
	UpscaleLimit         int     `json:"upscale_limit"`
	DownscaleLimit       int     `json:"downscale_limit"`
	ScaleUpOn503         *bool   `json:"scale_up_on_503"`
	NewRelicApiKey       *string `json:"new_relic_api_key"`
	NewRelicAccountId    *string `json:"new_relic_account_id"`
	NewRelicAppId        *string `json:"new_relic_app_id"`
	Notify               bool    `json:"notify"`
	NotifyQuantity       int     `json:"notify_quantity"`
	NotifyAfter          int     `json:"notify_after"`
	UpscaleOnInitialJob  *bool   `json:"upscale_on_initial_job"`
}

type wrappedManager struct {
	Manager Manager `json:"manager"`
}

func (r *ManagerResource) Get(id string) (*Manager, error) {
	wrapped := &wrappedManager{}
	err := r.client.getResource("managers", id, wrapped)
	return &wrapped.Manager, err
}

func (r *ManagerResource) Create(create Manager) (*Manager, error) {
	wrapped := &wrappedManager{Manager: create}
	err := r.client.createResource("managers", wrapped)
	return &wrapped.Manager, err
}

func (r *ManagerResource) Update(update Manager) (*Manager, error) {
	wrapped := &wrappedManager{Manager: update}
	err := r.client.updateResource("managers", update.Id, wrapped)
	return &wrapped.Manager, err
}

func (r *ManagerResource) Delete(id string) error {
	err := r.client.deleteResource("managers", id)
	return err
}
