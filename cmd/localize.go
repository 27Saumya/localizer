package cmd

import (
	"fmt"
	"os"

	"github.com/27Saumya/localizer/internal"
	"github.com/spf13/cobra"
)

var localizeCmd = &cobra.Command{
	Use:   "localize [base .lproj directory] [languages]",
	Short: "Adds localization for the languages provided",
	Long: 
`Adds localization for the languages (separated by a comma if there are multiple) provided with reference to a base .lproj directory
The base .lproj directory can be any of your preferred language.
Both the .lproj directory and the languages provided should be a valid iso code

For example:-
localizer localize en.lproj fr de es
`,

	Args: func(cmd *cobra.Command, args []string) error {
		argsCount := len(args)
		if argsCount < 2 {
			return fmt.Errorf("❌ expected atleast '2' arguments, found '%d'", argsCount)
		}
		sourcePath := args[0]
		return internal.HandlePath(sourcePath)
	},

	Run: func(cmd *cobra.Command, args []string) {
		sourcePath := args[0]
		if sourceInfo, err := os.Stat(sourcePath); os.IsNotExist(err) {
			panic(err)
		} else {
			if sourceInfo.IsDir() {
				internal.Localize(sourcePath, args)
			} else {
				panic("❌ expected a '.lproj' directory, found a file")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(localizeCmd)
}
