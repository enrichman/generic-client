package client

import (
	"context"
	"fmt"
)

const (
	namespaceListPath   = "/namespaces"
	namespaceCreatePath = "/namespaces"
)

type NamespaceService struct {
	client *Client
}

type Namespace struct {
	Name           string   `json:"name,omitempty"`
	Apps           []string `json:"apps,omitempty"`
	Configurations []string `json:"configurations,omitempty"`
}

func (s *NamespaceService) List(ctx context.Context) ([]Namespace, error) {
	resp, err := get(ctx, s.client, namespaceCreatePath, []Namespace{})
	if err != nil {
		return nil, err
	}
	return resp.Response, nil
}

func (s *NamespaceService) Create(ctx context.Context, name string) error {
	type namespaceCreateRequest struct {
		Name string `json:"name"`
	}

	payload := &namespaceCreateRequest{
		Name: name,
	}

	resp, err := post(ctx, s.client, namespaceCreatePath, payload, &StatusResponse{})
	if err != nil {
		return err
	}

	if resp.Response.Status == "ko" {
		// do something
		fmt.Println("wrooong")
	}

	return nil
}
