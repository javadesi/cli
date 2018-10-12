package v2action_test

import "code.cloudfoundry.org/cli/actor/v2action/v2actionfakes"

var _ = Describe("Router Group Actions", func() {
	var (
		actor            *Actor
		fakeRouterClient *v2actionfakes.fakeRouterClient
	)

	BeforeEach(func() {
		fakeRouterClient = new(v2actionfakes.fakeRouterClient)
		actor = NewActor(nil, nil, nil)
	})

	Describe("Router Group", func() {
		Describe("List", func() {

			It("should list all the router groups", func() {
			})
		})
	})
})
