package ui

import (
	"context"
	"fmt"
	"io"

	"github.com/enrichman/generic-client/pkg/client"
)

type userLister interface {
	List(ctx context.Context) ([]client.User, error)
}

func List(writer io.Writer, userLister userLister) error {
	namespaces, err := userLister.List(context.Background())
	if err != nil {
		return err
	}

	fmt.Fprintln(writer, "namespaceList: ", namespaces)
	return nil
}
