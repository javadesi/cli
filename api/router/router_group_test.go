package router_test

import (
	"net/http"

	. "code.cloudfoundry.org/cli/api/router"
	"code.cloudfoundry.org/cli/api/router/routererror"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = FDescribe("Router Groups", func() {
	Describe("GetRouterGroups", func() {
		var (
			client       *Client
			fakeConfig   Config
			routerGroups []RouterGroup
			executeErr   error
		)

		JustBeforeEach(func() {
			fakeConfig = NewTestConfig()
			client = NewTestRouterClient(fakeConfig)
			routerGroups, executeErr = client.GetRouterGroups()
		})

		When("the request fails", func() {
			BeforeEach(func() {
				response := `{"name":"UnauthorizedError","message":"Token is expired"}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/routing/v1/router_groups"),
						VerifyHeaderKV("Content-Type", "application/json"),
						RespondWith(http.StatusUnauthorized, response),
					))
			})

			It("returns the error", func() {
				Expect(executeErr).To(HaveOccurred())
				expectedErr := routererror.ErrorResponse{
					Message:    "Token is expired",
					StatusCode: 401,
					Name:       "UnauthorizedError",
				}
				Expect(executeErr).To(MatchError(expectedErr))
				Expect(routerGroups).To(BeEmpty())
			})
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				response := `[
					{
						"guid":"some-router-group-guid-1",
						"name":"default-tcp",
						"type":"tcp",
						"reservable_ports":"1024-1123"
					},
					{
						"guid":"some-router-group-guid-2",
						"name":"test-router-group",
						"type":"test-tcp",
						"reservable_ports":"1234-2345"
					}
				]`

				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/routing/v1/router_groups"),
						VerifyHeaderKV("Content-Type", "application/json"),
						RespondWith(http.StatusOK, response),
					))
			})

			It("returns the list of router groups and no errors", func() {
				Expect(executeErr).NotTo(HaveOccurred())
				Expect(routerGroups).To(ConsistOf(RouterGroup{
					GUID:            "some-router-group-guid-1",
					Name:            "default-tcp",
					Type:            "tcp",
					ReservablePorts: "1024-1123",
				}, RouterGroup{
					GUID:            "some-router-group-guid-2",
					Name:            "test-router-group",
					Type:            "test-tcp",
					ReservablePorts: "1234-2345",
				}))
			})
		})
	})
})
