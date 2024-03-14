package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/27Saumya/localizer/internal"
	"github.com/fatih/color"
)


var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check what version of localizer you are using",
	Long: "Check what version of localizer you are using",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Localizer version: %s\n", color.GreenString(internal.Version))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
