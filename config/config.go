package config

import (
	"fmt"
	"github.com/carwow/terraform-provider-hirefire/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	Client *client.Client
}

func Init(d *schema.ResourceData) (interface{}, error) {
	config := &Config{}

	if value, ok := d.GetOk("api_key"); ok {
		config.Client = client.New(value.(string))
	} else {
		return nil, fmt.Errorf("Hirefire API key is required")
	}

	return config, nil
}

func Client(meta interface{}) *client.Client {
	return meta.(*Config).Client
}
