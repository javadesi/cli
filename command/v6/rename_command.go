package v6

import (
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/translatableerror"
)

//go:generate counterfeiter . RenameActor

type RenameActor interface {
}

type RenameCommand struct {
	RequiredArgs    flag.AppRenameArgs `positional-args:"yes"`
	usage           interface{}        `usage:"CF_NAME rename APP_NAME NEW_APP_NAME"`
	relatedCommands interface{}        `related_commands:"apps, delete"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       RenameActor
}

func (cmd *RenameCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	// ccClient, uaaClient, err := shared.NewClients(config, ui, true)
	// if err != nil {
	// 	return err
	// }
	// cmd.Actor = v2action.NewActor(ccClient, uaaClient, config)

	return nil
}

func (RenameCommand) Execute(args []string) error {
	return translatableerror.UnrefactoredCommandError{}
}
