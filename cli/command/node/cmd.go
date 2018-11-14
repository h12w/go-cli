package node

import (
	"github.com/spf13/cobra"

	"github.com/storageos/go-cli/cli"
	"github.com/storageos/go-cli/cli/command"
)

// NewNodeCommand returns a cobra command for `node` subcommands
func NewNodeCommand(storageosCli *command.StorageOSCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Manage nodes",
		Args:  cli.NoArgs,
		RunE:  storageosCli.ShowHelp,
	}
	cmd.AddCommand(
		command.WithAlias(newListCommand(storageosCli), command.ListAliases...),
		command.WithAlias(newInspectCommand(storageosCli), command.InspectAliases...),
		command.WithAlias(newHealthCommand(storageosCli), command.HealthAliases...),
		newCordonCommand(storageosCli),
		newUncordonCommand(storageosCli),
		newConnectivityCommand(storageosCli),
		newDrainCommand(storageosCli),
		newUndrainCommand(storageosCli),
		command.WithAlias(newUpdateCommand(storageosCli), command.UpdateAliases...),
		command.WithAlias(newDeleteCommand(storageosCli), command.RemoveAliases...),
	)
	return cmd
}
