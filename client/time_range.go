package client

type TimeRangeResource struct {
	client *Client
}

type TimeRange struct {
	Id          string `json:"id"`
	ManagerId   string `json:"manager_id"`
	FromMinute  int    `json:"from_minute"`
	UntilMinute int    `json:"until_minute"`
	Minimum     int    `json:"minimum"`
	Maximum     int    `json:"maximum"`
	Position    int    `json:"position"`
	Monday      bool   `json:"monday"`
	Tuesday     bool   `json:"tuesday"`
	Wednesday   bool   `json:"wednesday"`
	Thursday    bool   `json:"thursday"`
	Friday      bool   `json:"friday"`
	Saturday    bool   `json:"saturday"`
	Sunday      bool   `json:"sunday"`
}

type wrappedTimeRange struct {
	TimeRange TimeRange `json:"time_range"`
}

func (r *TimeRangeResource) Get(id string) (*TimeRange, error) {
	wrapped := &wrappedTimeRange{}
	err := r.client.getResource("time_ranges", id, wrapped)
	return &wrapped.TimeRange, err
}

func (r *TimeRangeResource) Create(create TimeRange) (*TimeRange, error) {
	wrapped := &wrappedTimeRange{TimeRange: create}
	err := r.client.createResource("time_ranges", wrapped)
	return &wrapped.TimeRange, err
}

func (r *TimeRangeResource) Update(update TimeRange) (*TimeRange, error) {
	wrapped := &wrappedTimeRange{TimeRange: update}
	err := r.client.updateResource("time_ranges", update.Id, wrapped)
	return &wrapped.TimeRange, err
}

func (r *TimeRangeResource) Delete(id string) error {
	err := r.client.deleteResource("time_ranges", id)
	return err
}
