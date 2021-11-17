package release

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/flow-cli/internal/command"

	"github.com/onflow/flow-cli/internal/accounts"
	"github.com/onflow/flow-cli/internal/cadence"
	"github.com/onflow/flow-cli/internal/collections"
	"github.com/onflow/flow-cli/internal/config"
	"github.com/onflow/flow-cli/internal/emulator"
	"github.com/onflow/flow-cli/internal/events"
	"github.com/onflow/flow-cli/internal/project"
	"github.com/onflow/flow-cli/internal/scripts"
	"github.com/onflow/flow-cli/internal/transactions"
	"github.com/onflow/flow-cli/internal/version"
	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/blocks"
)

func TestChecklist(t *testing.T) {

	var cmd = &cobra.Command{
		Use:              "flow",
		TraverseChildren: true,
	}

	// hot Commands
	config.InitCommand.AddToParent(cmd)

	// structured Commands
	cmd.AddCommand(cadence.Cmd)
	cmd.AddCommand(version.Cmd)
	cmd.AddCommand(emulator.Cmd)
	cmd.AddCommand(accounts.Cmd)
	cmd.AddCommand(scripts.Cmd)
	cmd.AddCommand(transactions.Cmd)
	//cmd.AddCommand(keys.Cmd) removed
	cmd.AddCommand(events.Cmd)      // changed descs
	cmd.AddCommand(blocks.Cmd)      // add flag, removed flag
	cmd.AddCommand(collections.Cmd) // new properties
	cmd.AddCommand(project.Cmd)

	command.InitFlags(cmd)

	err := checkVersionWithCurrent(cmd, "6", "5")

	assert.NoError(t, err)
}
