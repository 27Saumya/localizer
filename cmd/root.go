package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "localizer",
	Short: "Quickly add localization in your Xcode application",
	Long: 
`Localizer is a CLI tool to quickly and effeciently add localization in your Xcode application.
As many languages available can be localized with accuracy not being an issue.
Plus the best part, no configuration is required to use it!`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}