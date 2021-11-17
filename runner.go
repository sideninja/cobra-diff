package release

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var gray = "\033[37m"
var white = "\033[97m"

func checkVersionWithCurrent(topCommand *cobra.Command, currentVersion string, previousVersion string) error {
	previous := Specs{Version: previousVersion}
	err := previous.load()
	if err != nil {
		return err
	}

	fmt.Printf("Version Compare from v%s -> v%s\n", previousVersion, currentVersion)
	current := generateSpecs(topCommand, currentVersion)

	removedDiffs := current.diff(previous)
	//addedChanges := previous.diff(current)

	for _, diff := range removedDiffs {
		if diff.usage == nil && diff.example == nil && diff.short == nil && diff.missing == false {
			fmt.Printf("✅ ")
		} else if diff.missing {
			fmt.Printf("❌ %s ", pRed("[removed]"))
		}
		fmt.Printf("%s\n", diff.path)

		if diff.usage != nil || diff.example != nil || diff.short != nil {
			fmt.Printf("\tProperties:\n")
		}

		if diff.usage != nil {
			fmt.Printf(
				"\t%s usage\n\t\t(v%s) %s\n\t\t(v%s) %s\n\n",
				pBlue("[changed]"),
				previousVersion,
				diff.usage.valueA(),
				currentVersion,
				diff.usage.valueB(),
			)
		}
		if diff.example != nil {
			fmt.Printf(
				"\t%s example\n\t\t(v%s) %s\n\t\t(v%s) %s\n\n",
				pBlue("[changed]"),
				previousVersion,
				pRed(diff.example.valueA()),
				currentVersion,
				pGreen(diff.example.valueB()),
			)
		}
		if diff.short != nil {
			fmt.Printf(
				"\t%s short\n\t\t(v%s) %s\n\t\t(v%s) %s\n\n",
				pBlue("[changed]"),
				previousVersion,
				pRed(diff.example.valueA()),
				currentVersion,
				pGreen(diff.example.valueB()),
			)
		}

		if len(diff.flags) > 0 {
			for _, fDiff := range diff.flags {
				if fDiff.usage != nil || fDiff.defaultValue != nil {
					fmt.Printf("\tFlags:\n")
				}

				if fDiff.usage != nil {
					fmt.Printf(
						"\t%s usage\n\t\t(v%s) %s\n\t\t(v%s) %s\n\n",
						pBlue("[changed]"),
						previousVersion,
						pRed(fDiff.usage.valueA()),
						currentVersion,
						pGreen(fDiff.usage.valueB()),
					)
				}
				if fDiff.defaultValue != nil {
					fmt.Printf(
						"\t%s default value\n\t\t(v%s) %s\n\t\t(v%s) %s\n\n",
						pBlue("[changed]"),
						previousVersion,
						pRed(fDiff.defaultValue.valueA()),
						currentVersion,
						pGreen(fDiff.defaultValue.valueB()),
					)
				}
			}
		}

		fmt.Printf("\n")
	}

	return nil
}

func pRed(v string) string {
	return fmt.Sprintf("%s%s%s", red, v, reset)
}
func pBlue(v string) string {
	return fmt.Sprintf("%s%s%s", blue, v, reset)
}
func pYellow(v string) string {
	return fmt.Sprintf("%s%s%s", yellow, v, reset)
}
func pGreen(v string) string {
	return fmt.Sprintf("%s%s%s", green, v, reset)
}
