package suite_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/enrichman/generic-client/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing the namespace APIs", Ordered, Label("users"), func() {
	var srv *httptest.Server
	var handler http.Handler
	var epinioClient *client.Client

	BeforeAll(func() {
		srv = httptest.NewServer(handler)

		epinioClient = client.NewClient(http.DefaultClient, srv.URL)
	})

	AfterAll(func() {
		srv.Close()
	})

	When("the listing the namespaces", func() {

		BeforeEach(func() {
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`[{"name":"ns-1"},{"name":"ns-2"}]`))
			})
		})

		Context("and the book is available", func() {
			It("lends it to the reader", func() {

				res, err := epinioClient.Users.List(context.Background())
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
			})
		})
	})
})
