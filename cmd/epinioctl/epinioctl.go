package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/enrichman/epinio-client-go/pkg/client"
)

func main() {
	ep := client.NewClient(http.DefaultClient, "http://localhost")
	res, err := ep.Namespace.List(context.Background())
	fmt.Printf("%+v - %+v", res, err)
}
