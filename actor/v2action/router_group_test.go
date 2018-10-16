package v2action_test

import (
	"errors"

	"code.cloudfoundry.org/cli/actor/actionerror"
	. "code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2action/v2actionfakes"
	"code.cloudfoundry.org/cli/api/router"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Router Group Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient = new(v2actionfakes.FakeCloudControllerClient)
		fakeRouterClient          *v2actionfakes.FakeRouterClient
	)

	BeforeEach(func() {
		fakeRouterClient = new(v2actionfakes.FakeRouterClient)
		actor = NewActor(fakeCloudControllerClient, nil, nil)
	})

	Describe("GetRouterGroupByName", func() {
		var (
			routerGroupName string
			routerGroup     RouterGroup
			err             error
		)

		JustBeforeEach(func() {
			routerGroup, err = actor.GetRouterGroupByName(routerGroupName, fakeRouterClient)
		})

		When("the router group does not exists", func() {
			BeforeEach(func() {
				routerGroupName = "some-router-group"
				fakeRouterClient.GetRouterGroupsReturns([]router.RouterGroup{
					router.RouterGroup{Name: "some-other-router-group"},
					router.RouterGroup{Name: "some-entirely-different-router-group"},
				}, nil)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(actionerror.RouterGroupNotFoundError{Name: routerGroupName}))
				Expect(routerGroup).To(Equal(RouterGroup{}))
				Expect(fakeRouterClient.GetRouterGroupsCallCount()).To(Equal(1))
			})
		})

		When("the router group exists", func() {
			BeforeEach(func() {
				routerGroupName = "default-tcp"
				fakeRouterClient.GetRouterGroupsReturns([]router.RouterGroup{router.RouterGroup{Name: routerGroupName}}, nil)
			})

			It("should return the router group and not an error", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(routerGroup).To(Equal(RouterGroup{Name: routerGroupName}))
				Expect(fakeRouterClient.GetRouterGroupsCallCount()).To(Equal(1))
			})
		})

		When("the router client returns an error", func() {
			BeforeEach(func() {
				routerGroupName = "default-tcp"
				fakeRouterClient.GetRouterGroupsReturns([]router.RouterGroup{}, errors.New("The request failed"))
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("The request failed"))
				Expect(fakeRouterClient.GetRouterGroupsCallCount()).To(Equal(1))
			})

		})
	})
})
