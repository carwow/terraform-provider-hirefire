package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	http   *http.Client
	URL    string
	apiKey string

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
		http:   &http.Client{},
		URL:    DefaultURL,
		apiKey: apiKey,
	}

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

func (c *Client) do(method, path string, in interface{}) (int, []byte, error) {
	var body io.Reader
	if in != nil {
		data, err := json.Marshal(in)
		if err != nil {
			return 0, nil, err
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.URL+path, body)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Accept", "application/vnd.hirefire.v1+json")
	req.Header.Set("Authorization", "Token "+c.apiKey)
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.http.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, out, nil
}

func (c *Client) request(method, path string, in interface{}, expectedStatus int, out interface{}) error {
	status, body, err := c.do(method, path, in)
	if err != nil {
		return err
	}
	if status != expectedStatus {
		return fmt.Errorf("%d: %s", status, body)
	}
	if out == nil {
		return nil
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("%s: %s", err, body)
	}
	return nil
}

func (c *Client) getResource(path string, id string, wrapped interface{}) error {
	if id == "" {
		return fmt.Errorf("ID can't be empty")
	}
	return c.request(http.MethodGet, path+"/"+id, nil, http.StatusOK, wrapped)
}

func (c *Client) createResource(path string, wrapped interface{}) error {
	return c.request(http.MethodPost, path, wrapped, http.StatusCreated, wrapped)
}

func (c *Client) updateResource(path string, id string, wrapped interface{}) error {
	if id == "" {
		return fmt.Errorf("ID can't be empty")
	}
	return c.request(http.MethodPatch, path+"/"+id, wrapped, http.StatusOK, wrapped)
}

func (c *Client) deleteResource(path string, id string) error {
	if id == "" {
		return fmt.Errorf("ID can't be empty")
	}
	return c.request(http.MethodDelete, path+"/"+id, nil, http.StatusOK, nil)
}
