package gencli

import (
	"context"
)

type UserService interface {
	List(ctx context.Context) ([]User, error)
	Create(ctx context.Context, name string) (User, error)
}

type User struct {
	Name string
}

func CreateUser(ctx context.Context, userService UserService, name string) (User, error) {
	return userService.Create(ctx, name)
}
