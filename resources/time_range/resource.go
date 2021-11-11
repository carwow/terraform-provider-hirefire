package time_range

import (
	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/carwow/terraform-provider-hirefire/config"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"manager_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"from_minute": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"until_minute": &schema.Schema{
				Type:     schema.TypeInt,
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
			"position": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monday": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tuesday": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"wednesday": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"thursday": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"friday": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"saturday": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sunday": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func setAttributes(d *schema.ResourceData, timeRange *client.TimeRange) {
	d.Set("manager_id", timeRange.ManagerId)
	d.Set("from_minute", timeRange.FromMinute)
	d.Set("until_minute", timeRange.UntilMinute)
	d.Set("minimum", timeRange.Minimum)
	d.Set("maximum", timeRange.Maximum)
	d.Set("position", timeRange.Position)
	d.Set("monday", timeRange.Monday)
	d.Set("tuesday", timeRange.Tuesday)
	d.Set("wednesday", timeRange.Wednesday)
	d.Set("thursday", timeRange.Thursday)
	d.Set("friday", timeRange.Friday)
	d.Set("saturday", timeRange.Saturday)
	d.Set("sunday", timeRange.Sunday)
}

func getAttributes(d *schema.ResourceData) client.TimeRange {
	return client.TimeRange{
		Id:          d.Id(),
		ManagerId:   d.Get("manager_id").(string),
		FromMinute:  d.Get("from_minute").(int),
		UntilMinute: d.Get("until_minute").(int),
		Minimum:     d.Get("minimum").(int),
		Maximum:     d.Get("maximum").(int),
		Position:    d.Get("position").(int),
		Monday:      d.Get("monday").(bool),
		Tuesday:     d.Get("tuesday").(bool),
		Wednesday:   d.Get("wednesday").(bool),
		Thursday:    d.Get("thursday").(bool),
		Friday:      d.Get("friday").(bool),
		Saturday:    d.Get("saturday").(bool),
		Sunday:      d.Get("sunday").(bool),
	}
}

func create(d *schema.ResourceData, m interface{}) error {
	timeRange, err := config.Client(m).TimeRange.Create(getAttributes(d))
	if err != nil {
		return err
	}

	d.SetId(timeRange.Id)
	setAttributes(d, timeRange)
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	_, err := config.Client(m).TimeRange.Update(getAttributes(d))
	if err != nil {
		return err
	}

	d.Partial(false)
	return nil
}

func delete(d *schema.ResourceData, m interface{}) error {
	err := config.Client(m).TimeRange.Delete(d.Id())
	return err
}

func importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	err := read(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
