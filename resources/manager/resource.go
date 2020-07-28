package manager

import (
	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/config"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: delete,

		Importer: &schema.ResourceImporter{
			State: importer,
		},

		Schema: map[string]*schema.Schema{
			"application_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Manager::Worker::HireFire::JobQueue",
					"Manager::Web::Logplex::Load",
					"Manager::Web::Logplex::ResponseTime",
					"Manager::Web::Logplex::ConnectTime",
					"Manager::Web::Logplex::QueueTime",
					"Manager::Web::Logplex::RPM",
					"Manager::Web::NewRelic::V2::ResponseTime",
					"Manager::Web::NewRelic::V2::RPM",
					"Manager::Web::NewRelic::V2::Apdex",
					"Manager::Web::NewRelic::V1::ResponseTime",
					"Manager::Web::NewRelic::V1::RPM",
					"Manager::Web::NewRelic::V1::Apdex",
					"Manager::Web::HireFire::ResponseTime",
				}, false),
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"maximum": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"aggregation": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"average", "percentile"}, false),
			},
			"percentile": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_queue_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_queue_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_response_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_response_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_apdex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_apdex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_minutes": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 5, 15}),
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"decrementable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upscale_quantity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"downscale_quantity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upscale_sensitivity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"downscale_sensitivity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upscale_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"downscale_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upscale_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"downscale_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_up_on_503": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"new_relic_api_key": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"new_relic_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"new_relic_app_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"notify_quantity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"notify_after": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func setAttributes(d *schema.ResourceData, manager *client.Manager) {
	d.Set("application_id", manager.ApplicationId)
	d.Set("name", manager.Name)
	d.Set("type", manager.Type)
	d.Set("enabled", manager.Enabled)
	d.Set("minimum", manager.Minimum)
	d.Set("maximum", manager.Maximum)

	d.Set("aggregation", manager.Aggregation)
	d.Set("percentile", manager.Percentile)
	d.Set("minimum_queue_time", manager.MinimumQueueTime)
	d.Set("maximum_queue_time", manager.MaximumQueueTime)
	d.Set("minimum_response_time", manager.MinimumResponseTime)
	d.Set("maximum_response_time", manager.MaximumResponseTime)
	d.Set("minimum_load", manager.MinimumLoad)
	d.Set("maximum_load", manager.MaximumLoad)
	d.Set("minimum_apdex", manager.MinimumApdex)
	d.Set("maximum_apdex", manager.MaximumApdex)
	d.Set("last_minutes", manager.LastMinutes)
	d.Set("ratio", manager.Ratio)
	d.Set("decrementable", manager.Decrementable)
	d.Set("url", manager.Url)
	d.Set("upscale_quantity", manager.UpscaleQuantity)
	d.Set("downscale_quantity", manager.DownscaleQuantity)
	d.Set("upscale_sensitivity", manager.UpscaleSensitivity)
	d.Set("downscale_sensitivity", manager.DownscaleSensitivity)
	d.Set("upscale_timeout", manager.UpscaleTimeout)
	d.Set("downscale_timeout", manager.DownscaleTimeout)
	d.Set("upscale_limit", manager.UpscaleLimit)
	d.Set("downscale_limit", manager.DownscaleLimit)
	d.Set("scale_up_on_503", manager.ScaleUpOn503)
	d.Set("new_relic_api_key", manager.NewRelicApiKey)
	d.Set("new_relic_account_id", manager.NewRelicAccountId)
	d.Set("new_relic_app_id", manager.NewRelicAppId)
	d.Set("notify", manager.Notify)
	d.Set("notify_quantity", manager.NotifyQuantity)
	d.Set("notify_after", manager.NotifyAfter)
}

func getAttributes(d *schema.ResourceData) client.Manager {
	manager := client.Manager{
		Id:            d.Id(),
		ApplicationId: d.Get("application_id").(string),
		Name:          d.Get("name").(string),
		Type:          d.Get("type").(string),
		Enabled:       d.Get("enabled").(bool),
		Minimum:       d.Get("minimum").(int),
		Maximum:       d.Get("maximum").(int),
	}

	if v, ok := d.GetOk("aggregation"); ok {
		value := v.(string)
		manager.Aggregation = &value
	}

	if v, ok := d.GetOk("percentile"); ok {
		value := v.(int)
		manager.Percentile = &value
	}

	if v, ok := d.GetOk("minimum_queue_time"); ok {
		value := v.(int)
		manager.MinimumQueueTime = &value
	}

	if v, ok := d.GetOk("maximum_queue_time"); ok {
		value := v.(int)
		manager.MaximumQueueTime = &value
	}

	if v, ok := d.GetOk("minimum_response_time"); ok {
		value := v.(int)
		manager.MinimumResponseTime = &value
	}

	if v, ok := d.GetOk("maximum_response_time"); ok {
		value := v.(int)
		manager.MaximumResponseTime = &value
	}

	if v, ok := d.GetOk("minimum_load"); ok {
		value := v.(int)
		manager.MinimumLoad = &value
	}

	if v, ok := d.GetOk("maximum_load"); ok {
		value := v.(int)
		manager.MaximumLoad = &value
	}

	if v, ok := d.GetOk("minimum_apdex"); ok {
		value := v.(int)
		manager.MinimumApdex = &value
	}

	if v, ok := d.GetOk("maximum_apdex"); ok {
		value := v.(int)
		manager.MaximumApdex = &value
	}

	if v, ok := d.GetOk("last_minutes"); ok {
		value := v.(int)
		manager.LastMinutes = &value
	}

	if v, ok := d.GetOk("ratio"); ok {
		value := v.(int)
		manager.Ratio = &value
	}

	if v, ok := d.GetOk("decrementable"); ok {
		value := v.(bool)
		manager.Decrementable = &value
	}

	if v, ok := d.GetOk("url"); ok {
		value := v.(string)
		manager.Url = &value
	}

	if v, ok := d.GetOk("upscale_quantity"); ok {
		value := v.(int)
		manager.UpscaleQuantity = &value
	}

	if v, ok := d.GetOk("downscale_quantity"); ok {
		value := v.(int)
		manager.DownscaleQuantity = &value
	}

	if v, ok := d.GetOk("upscale_sensitivity"); ok {
		value := v.(int)
		manager.UpscaleSensitivity = &value
	}

	if v, ok := d.GetOk("downscale_sensitivity"); ok {
		value := v.(int)
		manager.DownscaleSensitivity = &value
	}

	if v, ok := d.GetOk("upscale_timeout"); ok {
		value := v.(int)
		manager.UpscaleTimeout = &value
	}

	if v, ok := d.GetOk("downscale_timeout"); ok {
		value := v.(int)
		manager.DownscaleTimeout = &value
	}

	if v, ok := d.GetOk("upscale_limit"); ok {
		value := v.(int)
		manager.UpscaleLimit = &value
	}

	if v, ok := d.GetOk("downscale_limit"); ok {
		value := v.(int)
		manager.DownscaleLimit = &value
	}

	if v, ok := d.GetOk("scale_up_on_503"); ok {
		value := v.(bool)
		manager.ScaleUpOn503 = &value
	}

	if v, ok := d.GetOk("new_relic_api_key"); ok {
		value := v.(string)
		manager.NewRelicApiKey = &value
	}

	if v, ok := d.GetOk("new_relic_account_id"); ok {
		value := v.(string)
		manager.NewRelicAccountId = &value
	}

	if v, ok := d.GetOk("new_relic_app_id"); ok {
		value := v.(string)
		manager.NewRelicAppId = &value
	}

	if v, ok := d.GetOk("notify"); ok {
		value := v.(bool)
		manager.Notify = &value
	}

	if v, ok := d.GetOk("notify_quantity"); ok {
		value := v.(int)
		manager.NotifyQuantity = &value
	}

	if v, ok := d.GetOk("notify_after"); ok {
		value := v.(int)
		manager.NotifyAfter = &value
	}

	return manager
}

func create(d *schema.ResourceData, m interface{}) error {
	manager, err := config.Client(m).Manager.Create(getAttributes(d))
	if err != nil {
		return err
	}

	d.SetId(manager.Id)
	setAttributes(d, manager)
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	manager, err := config.Client(m).Manager.Get(d.Id())
	if err != nil {
		return err
	}

	setAttributes(d, manager)
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	_, err := config.Client(m).Manager.Update(getAttributes(d))
	if err != nil {
		return err
	}

	d.Partial(false)
	return nil
}

func delete(d *schema.ResourceData, m interface{}) error {
	err := config.Client(m).Manager.Delete(d.Id())
	return err
}

func importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	err := read(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
