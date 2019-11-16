package client

import (
	"fmt"
	"github.com/imroc/req"
)

type Client struct {
	req    *req.Req
	url    string
	apiKey string

	Organization *OrganizationResource
	Account      *AccountResource
}

const DefaultURL = "https://api.hirefire.io/"

func New(apiKey string) *Client {
	client := &Client{
		req:    req.New(),
		url:    DefaultURL,
		apiKey: apiKey,
	}

	client.Organization = &OrganizationResource{client: client}
	client.Account = &AccountResource{client: client}

	return client
}

func (c *Client) get(path string, v ...interface{}) (*req.Resp, error) {
	defaults := []interface{}{
		req.Header{
			"Accept":        "application/vnd.hirefire.v1+json",
			"Authorization": "Token " + c.apiKey,
		},
	}
	return c.req.Get(c.url+path, append(defaults, v...)...)
}

func (c *Client) getResource(path string, id string, wrapped interface{}) error {
	res, err := c.get(path + "/" + id)
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 200 {
		return fmt.Errorf("%s", res.String())
	}

	err = res.ToJSON(&wrapped)
	if err != nil {
		return fmt.Errorf("%s: %s", err, res.String())
	}

	return nil
}
