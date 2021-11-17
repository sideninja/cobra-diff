package release

import (
	"github.com/spf13/pflag"
)

type flagsSpecs []flagSpec

func (f *flagsSpecs) diff(refFlags flagsSpecs) []flagDiff {
	diffs := make([]flagDiff, 0)

	for _, refFlag := range refFlags {
		flag := f.find(refFlag)

		if flag != nil { // if flag is found compare it
			diffs = append(diffs, refFlag.diff(*flag))
		} else { // if properties not found store nil
			//diffs[refFlag.Name] = []Diff{{missing: true}} todo missing
		}
	}

	return diffs
}

func (f flagsSpecs) find(findFlag flagSpec) *flagSpec {
	for _, flag := range f {
		if flag.matches(findFlag) {
			return &flag
		}
	}
	return nil
}

type flagSpec struct {
	Name         string
	Usage        string
	DefaultValue string
}

func (f flagSpec) matches(spec flagSpec) bool {
	return f.Name == spec.Name
}

func (f flagSpec) diff(spec flagSpec) flagDiff {
	diff := newFlagDiff(f.Name)

	if f.Usage != spec.Usage {
		diff.usage = newStringDiff(f.Usage, spec.Usage)
	}
	if f.DefaultValue != spec.DefaultValue {
		diff.defaultValue = newStringDiff(f.DefaultValue, spec.DefaultValue)
	}

	return diff
}

func newFlags(flags *pflag.FlagSet) flagsSpecs {
	var flagsSpec flagsSpecs

	flags.VisitAll(func(flag *pflag.Flag) {
		f := flagSpec{
			Name:         flag.Name,
			DefaultValue: flag.DefValue,
			Usage:        flag.Usage,
		}
		flagsSpec = append(flagsSpec, f)
	})

	return flagsSpec
}
