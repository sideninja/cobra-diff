/*
 * Flow CLI
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
