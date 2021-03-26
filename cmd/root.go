package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "go-gsutil",
}

func Execute() error {
	return rootCmd.Execute()
}
