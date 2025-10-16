package api_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	"github.com/steinfletcher/apitest"
)

var _ = Describe("Boards", func() {
	var (
		t GinkgoTInterface
	)

	BeforeEach(func() {
		t = GinkgoT()
	})

	When("GET /boards", func() {
		It("should return boards", func() {
			handler := func(w http.ResponseWriter, r *http.Request) {
				msg := `{"message": "hello"}`
				_, _ = w.Write([]byte(msg))
				w.WriteHeader(http.StatusOK)
			}

			apitest.New(). // configuration
				HandlerFunc(handler).
				Get("/message"). // request
				Expect(t). // expectations
				Body(`{"message": "hello"}`).
				Status(http.StatusOK).
				End()
		})
	})
})
