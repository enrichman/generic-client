package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	BaseURL    string

	Namespace *NamespaceService
}

type GenericResponse[T any] struct {
	StatusCode int
	Response   T
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewClient(httpClient *http.Client, baseURL string) *Client {
	c := &Client{
		httpClient: httpClient,
		BaseURL:    baseURL,
	}

	c.Namespace = &NamespaceService{c}

	return c
}

func get[R any](ctx context.Context, c *Client, path string, response R) (GenericResponse[R], error) {
	return do(ctx, c, http.MethodGet, path, nil, response)
}

func post[R any](ctx context.Context, c *Client, path string, payload any, response R) (GenericResponse[R], error) {
	return do(ctx, c, http.MethodGet, path, nil, response)
}

func do[R any](ctx context.Context, c *Client, method, path string, payload any, response R) (GenericResponse[R], error) {
	url := c.BaseURL + path

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return GenericResponse[R]{}, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return GenericResponse[R]{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return GenericResponse[R]{}, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GenericResponse[R]{}, err
	}

	fmt.Println(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return GenericResponse[R]{}, err
	}

	//err = json.NewDecoder(res.Body).Decode(&response)

	return GenericResponse[R]{
		Response:   response,
		StatusCode: res.StatusCode,
	}, nil
}
