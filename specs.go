package release

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

func generateSpecs(topCommand *cobra.Command, version string) Specs {
	command := make(commandsSpecs, 0)

	specs := Specs{
		Commands: command,
		Version:  version,
	}
	specs.generate(topCommand)

	return specs
}

type Specs struct {
	Version  string
	Commands commandsSpecs
}

func (s *Specs) generate(cmd *cobra.Command) {
	c := newCommand(cmd)
	s.Commands = append(s.Commands, c)

	if cmd.HasSubCommands() {
		for _, sub := range cmd.Commands() {
			s.generate(sub)
		}
	}
}

func (s *Specs) save() error {
	specs, err := json.MarshalIndent(s.Commands, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(
		fmt.Sprintf("%s.json", s.Version),
		specs,
		0644,
	)
}

func (s *Specs) load() error {
	loadedSpecs, err := ioutil.ReadFile(
		fmt.Sprintf("%s.json", s.Version),
	)
	if err != nil {
		return err
	}

	var specs commandsSpecs
	err = json.Unmarshal(loadedSpecs, &specs)
	if err != nil {
		return err
	}
	s.Commands = specs

	return nil
}

func (s *Specs) validate() {
	// todo check if specs contain all necessary things (example, desc, not long desc, docs...)
}

func (s *Specs) documentationExists() {

}

func (s *Specs) diff(refSpec Specs) []commandDiff {
	return s.Commands.diff(refSpec.Commands)
}
