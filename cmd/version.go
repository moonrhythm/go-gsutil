package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.0-dev"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info about go-gsutil",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-gsutil version:", version)
	},
}
