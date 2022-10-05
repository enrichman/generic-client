package client

import (
	"context"

	"github.com/enrichman/generic-client/internal/gencli"
)

type UserService struct {
	client *Client
}

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

func (s *UserService) List(ctx context.Context) ([]gencli.User, error) {
	userResp, err := get(ctx, s.client, "/users", []UserResponse{})
	return makeUserList(userResp), err
}

func (s *UserService) GetUserByID(ctx context.Context, ID string) (gencli.User, error) {
	userResp, err := get(ctx, s.client, "/users/"+ID, UserResponse{})
	return makeUser(userResp), err
}

func (s *UserService) Create(ctx context.Context, name string) (gencli.User, error) {
	type userCreateRequest struct {
		Name string `json:"name"`
	}

	payload := &userCreateRequest{
		Name: name,
	}

	userResp, err := post(ctx, s.client, "/users", payload, UserResponse{})
	return makeUser(userResp), err
}

func makeUserList(userListResp []UserResponse) []gencli.User {
	userList := []gencli.User{}
	for _, userResp := range userListResp {
		userList = append(userList, makeUser(userResp))
	}
	return userList
}

func makeUser(userResp UserResponse) gencli.User {
	return gencli.User{
		Name: userResp.Name,
	}
}
