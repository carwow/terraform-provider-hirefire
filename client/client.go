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
	TimeRange    *TimeRangeResource
	Recipient    *RecipientResource
	User         *UserResource
	Membership   *MembershipResource
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

	// Uncomment to dump all HTTP requests and responses
	// req.Debug = true

	client.Organization = &OrganizationResource{client: client}
	client.Account = &AccountResource{client: client}
	client.Application = &ApplicationResource{client: client}
	client.Manager = &ManagerResource{client: client}
	client.TimeRange = &TimeRangeResource{client: client}
	client.Recipient = &RecipientResource{client: client}
	client.User = &UserResource{client: client}
	client.Membership = &MembershipResource{client: client}

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
	if id == "" {
		return fmt.Errorf("ID can't be empty")
	}
	res, err := c.get(path + "/" + id)
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 200 {
		return fmt.Errorf("%d: %s", res.Response().StatusCode, res.String())
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
		return fmt.Errorf("%d: %s", res.Response().StatusCode, res.String())
		// Uncomment to dump full HTTP request and response with error message
		// return fmt.Errorf("%d: %s\n%s", res.Response().StatusCode, res.String(), res.Dump())
	}

	err = res.ToJSON(&wrapped)
	if err != nil {
		return fmt.Errorf("%s: %s", err, res.String())
	}

	return nil
}

func (c *Client) updateResource(path string, id string, wrapped interface{}) error {
	if id == "" {
		return fmt.Errorf("ID can't be empty")
	}
	res, err := c.update(path+"/"+id, req.BodyJSON(&wrapped))
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 200 {
		return fmt.Errorf("%d: %s", res.Response().StatusCode, res.String())
	}

	err = res.ToJSON(&wrapped)
	if err != nil {
		return fmt.Errorf("%s: %s", err, res.String())
	}

	return nil
}

func (c *Client) deleteResource(path string, id string) error {
	if id == "" {
		return fmt.Errorf("ID can't be empty")
	}
	res, err := c.delete(path + "/" + id)
	if err != nil {
		return err
	}
	if res.Response().StatusCode != 200 {
		return fmt.Errorf("%d: %s", res.Response().StatusCode, res.String())
	}

	return nil
}
