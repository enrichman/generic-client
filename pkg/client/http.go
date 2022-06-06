package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	BaseURL    string

	Users *UserService
}

func NewClient(httpClient *http.Client, baseURL string) *Client {
	c := &Client{
		httpClient: httpClient,
		BaseURL:    baseURL,
	}

	c.Users = &UserService{c}

	return c
}

func get[R any](ctx context.Context, c *Client, path string, response R) (R, error) {
	return do(ctx, c, http.MethodGet, path, nil, response)
}

func post[R any](ctx context.Context, c *Client, path string, payload any, response R) (R, error) {
	return do(ctx, c, http.MethodPost, path, nil, response)
}

func do[R any](ctx context.Context, c *Client, method, path string, payload any, response R) (R, error) {
	emptyData := new(R)

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return *emptyData, err
		}
	}

	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return *emptyData, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return *emptyData, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return *emptyData, err
	}

	return response, nil
}
