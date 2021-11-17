package release

import (
	"github.com/spf13/cobra"
)

type commandSpec struct {
	Path           string
	Name           string
	Flags          flagsSpecs
	InheritedFlags flagsSpecs
	Example        string
	Usage          string
	Short          string
}

func (c commandSpec) diff(refCmd commandSpec) commandDiff {
	diff := newCommandDiff(c.Path)

	if c.Usage != refCmd.Usage {
		diff.usage = newStringDiff(c.Usage, refCmd.Usage)
	}
	if c.Example != refCmd.Example {
		diff.example = newStringDiff(c.Example, refCmd.Example)
	}
	if c.Short != refCmd.Short {
		diff.short = newStringDiff(c.Short, refCmd.Short)
	}

	diff.flags = c.Flags.diff(refCmd.Flags)
	diff.flags = append(diff.flags, c.InheritedFlags.diff(refCmd.InheritedFlags)...)

	return diff
}

func (c commandSpec) matches(spec commandSpec) bool {
	return c.Path == spec.Path
}

func newCommand(cmd *cobra.Command) commandSpec {
	return commandSpec{
		Path:           cmd.CommandPath(),
		Name:           cmd.Name(),
		InheritedFlags: newFlags(cmd.InheritedFlags()),
		Flags:          newFlags(cmd.NonInheritedFlags()),
		Example:        cmd.Example,
		Usage:          cmd.Use,
		Short:          cmd.Short,
	}
}

type commandsSpecs []commandSpec

func (c commandsSpecs) diff(refsCmd commandsSpecs) []commandDiff {
	diffs := make([]commandDiff, 0)

	for _, refCmd := range refsCmd {
		cmd := c.find(refCmd)

		if cmd != nil { // if properties is found compare it
			diffs = append(diffs, refCmd.diff(*cmd))
		} else { // if properties not found store missing
			diffs = append(diffs, newCommandMissing(refCmd.Path))
		}
	}

	return diffs
}

func (c commandsSpecs) find(spec commandSpec) *commandSpec {
	for _, cmd := range c {
		if cmd.matches(spec) {
			return &cmd
		}
	}

	return nil
}
