package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enrichman/generic-client/pkg/client"
	// . "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

func TestAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(listUsersHandler))
	defer srv.Close()

	gc := client.NewClient(http.DefaultClient, srv.URL)
	users, err := gc.Users.List(context.Background())
	fmt.Println(users, err)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`[{"name":"ns-1"},{"name":"ns-2"}]`))
}
