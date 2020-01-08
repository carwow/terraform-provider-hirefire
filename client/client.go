package client

import (
	"fmt"
	"github.com/imroc/req"
)

type Client struct {
	req     *req.Req
	URL     string
	headers []interface{}

	Organization *OrganizationResource
	Account      *AccountResource
	Application  *ApplicationResource
	Manager      *ManagerResource
}

const DefaultURL = "https://api.hirefire.io/"

func New(apiKey string) *Client {
	client := &Client{
		req: req.New(),
		URL: DefaultURL,
		headers: []interface{}{
			req.Header{
				"Accept":        "application/vnd.hirefire.v1+json",
				"Authorization": "Token " + apiKey,
			},
		},
	}

	client.Organization = &OrganizationResource{client: client}
	client.Account = &AccountResource{client: client}
	client.Application = &ApplicationResource{client: client}
	client.Manager = &ManagerResource{client: client}

	return client
}

func (c *Client) get(path string, v ...interface{}) (*req.Resp, error) {
	return c.req.Get(c.URL+path, append(c.headers, v...)...)
}

func (c *Client) create(path string, v ...interface{}) (*req.Resp, error) {
	return c.req.Post(c.URL+path, append(c.headers, v...)...)
}

func (c *Client) update(path string, v ...interface{}) (*req.Resp, error) {
	return c.req.Patch(c.URL+path, append(c.headers, v...)...)
}

func (c *Client) delete(path string, v ...interface{}) (*req.Resp, error) {
	return c.req.Delete(c.URL+path, append(c.headers, v...)...)
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

func (c *Client) createResource(path string, wrapped interface{}) error {
	res, err := c.create(path, req.BodyJSON(&wrapped))
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 201 {
		return fmt.Errorf("%s", res.String())
	}

	err = res.ToJSON(&wrapped)
	if err != nil {
		return fmt.Errorf("%s: %s", err, res.String())
	}

	return nil
}

func (c *Client) updateResource(path string, id string, wrapped interface{}) error {
	res, err := c.update(path+"/"+id, req.BodyJSON(&wrapped))
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

func (c *Client) deleteResource(path string, id string) error {
	res, err := c.delete(path + "/" + id)
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 200 {
		return fmt.Errorf("%s", res.String())
	}

	return nil
}
