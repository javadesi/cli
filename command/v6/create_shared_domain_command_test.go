package v6_test

import (
	"errors"

	"code.cloudfoundry.org/cli/actor/v2action"

	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	. "code.cloudfoundry.org/cli/command/v6"
	"code.cloudfoundry.org/cli/command/v6/v6fakes"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = FDescribe("CreateSharedDomainCommand", func() {
	var (
		fakeConfig      *commandfakes.FakeConfig
		fakeActor       *v6fakes.FakeCreateSharedDomainActor
		fakeSharedActor *commandfakes.FakeSharedActor
		testUI          *ui.UI
		cmd             CreateSharedDomainCommand

		executeErr       error
		sharedDomainName string
		username         string
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(v6fakes.FakeCreateSharedDomainActor)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		sharedDomainName = "some-shared-domain-name"

		cmd = CreateSharedDomainCommand{
			UI:           testUI,
			Config:       fakeConfig,
			Actor:        fakeActor,
			SharedActor:  fakeSharedActor,
			RequiredArgs: flag.Domain{Domain: sharedDomainName},
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	It("checks for user being logged in", func() {
		Expect(fakeSharedActor.RequireCurrentUserCallCount()).To(Equal(1))
	})

	When("user is logged in", func() {
		BeforeEach(func() {
			username = "some-user-name"
			fakeSharedActor.RequireCurrentUserReturns(username, nil)
		})

		When("the user is logged in as an admin", func() {
			When("--router-group is passed", func() {
				When("when the router group does not exists", func() {
					var expectedErr error

					BeforeEach(func() {
						expectedErr = errors.New("Not a real router group")
						fakeActor.GetRouterGroupByNameReturns(v2action.RouterGroup{}, expectedErr)
					})

					It("should fail and return error", func() {
						Expect(testUI.Out).To(Say("Creating shared domain %s as %s...", sharedDomainName, username))
						Expect(executeErr).To(MatchError(expectedErr))
					})
				})
			})
		})

		When("the user is not logged in as an admin", func() {

		})
	})

	When("the user is not logger in", func() {
		expectedErr := errors.New("not logged in and/or can't verify login because of error")

		BeforeEach(func() {
			fakeSharedActor.RequireCurrentUserReturns("", expectedErr)
		})

		It("returns the error", func() {
			Expect(executeErr).To(MatchError(expectedErr))
		})
	})

})
