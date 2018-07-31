package node

import (
	"context"
	"fmt"
	"strings"

	"github.com/dnephin/cobra"
	api "github.com/storageos/go-api"
	"github.com/storageos/go-api/types"
	"github.com/storageos/go-cli/cli"
	"github.com/storageos/go-cli/cli/command"
)

type leaveOptions struct {
	nodes []string
}

func newLeaveCommand(storageosCli *command.StorageOSCli) *cobra.Command {
	var opt leaveOptions

	cmd := &cobra.Command{
		Use:   "leave NODE [NODE...]",
		Short: "Remove one or more nodes from the cluster",
		Args:  cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opt.nodes = args
			return runLeave(storageosCli, opt)
		},
	}

	return cmd
}

func nodeLeave(client *api.Client, id string) error {
	return client.NodeDelete(types.DeleteOptions{
		ID:      id,
		Context: context.Background(),
	})
}

func runLeave(storageosCli *command.StorageOSCli, opt leaveOptions) error {
	client := storageosCli.Client()
	failed := make([]string, 0, len(opt.nodes))

	for _, id := range opt.nodes {
		if err := nodeLeave(client, id); err != nil {
			failed = append(failed, err.Error())
		}
	}

	return fmt.Errorf("failed to remove nodes: %s", strings.Join(failed, "\n"))
}
