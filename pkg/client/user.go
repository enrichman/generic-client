package client

import (
	"context"
)

type UserService struct {
	client *Client
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

func (s *UserService) List(ctx context.Context) ([]User, error) {
	return get(ctx, s.client, "/users", []User{})
}

func (s *UserService) GetUserByID(ctx context.Context, ID string) (User, error) {
	return get(ctx, s.client, "/users/"+ID, User{})
}

func (s *UserService) Create(ctx context.Context, name string) (User, error) {
	type userCreateRequest struct {
		Name string `json:"name"`
	}

	payload := &userCreateRequest{
		Name: name,
	}

	return post(ctx, s.client, "/users", payload, User{})
}
