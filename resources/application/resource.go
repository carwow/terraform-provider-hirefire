package application

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
			"account_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logplex_drain_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restart_crashed_dynos": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"new_issue_notifications": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"resolved_issue_notifications": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"checkup_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func setAttributes(d *schema.ResourceData, app *client.Application) {
	d.Set("account_id", app.AccountId)
	d.Set("name", app.Name)
	d.Set("custom_domain", app.CustomDomain)
	d.Set("logplex_drain_token", app.LogplexDrainToken)
	d.Set("ssl", app.Ssl)
	d.Set("restart_crashed_dynos", app.RestartCrashedDynos)
	d.Set("new_issue_notifications", app.NewIssueNotifications)
	d.Set("resolved_issue_notifications", app.ResolvedIssueNotifications)
	d.Set("token", app.Token)
	d.Set("checkup_frequency", app.CheckupFrequency)
}

func getAttributes(d *schema.ResourceData) client.Application {
	app := client.Application{
		Id:        d.Id(),
		AccountId: d.Get("account_id").(string),
		Name:      d.Get("name").(string),
	}

	if v, ok := d.GetOk("custom_domain"); ok {
		value := v.(string)
		app.CustomDomain = &value
	}

	if v, ok := d.GetOk("logplex_drain_token"); ok {
		value := v.(string)
		app.LogplexDrainToken = &value
	}

	if v, ok := d.GetOk("ssl"); ok {
		value := v.(bool)
		app.Ssl = value
	}

	if v, ok := d.GetOk("restart_crashed_dynos"); ok {
		value := v.(bool)
		app.RestartCrashedDynos = value
	}

	if v, ok := d.GetOk("new_issue_notifications"); ok {
		value := v.(bool)
		app.NewIssueNotifications = value
	}

	if v, ok := d.GetOk("resolved_issue_notifications"); ok {
		value := v.(bool)
		app.ResolvedIssueNotifications = value
	}

	if v, ok := d.GetOk("token"); ok {
		value := v.(string)
		app.Token = value
	}

	if v, ok := d.GetOk("checkup_frequency"); ok {
		value := v.(int)
		app.CheckupFrequency = value
	}

	return app
}

func create(d *schema.ResourceData, m interface{}) error {
	app, err := config.Client(m).Application.Create(getAttributes(d))
	if err != nil {
		return err
	}

	d.SetId(app.Id)
	setAttributes(d, app)
	return nil
}

func read(d *schema.ResourceData, m interface{}) error {
	app, err := config.Client(m).Application.Get(d.Id())
	if err != nil {
		return err
	}

	setAttributes(d, app)
	return nil
}

func update(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	_, err := config.Client(m).Application.Update(getAttributes(d))
	if err != nil {
		return err
	}

	d.Partial(false)
	return nil
}

func delete(d *schema.ResourceData, m interface{}) error {
	err := config.Client(m).Application.Delete(d.Id())
	return err
}

func importer(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	err := read(d, m)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
