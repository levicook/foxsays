package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "foxsays",
	Short: "foxsays is a twitter clone",
	Long:  ``,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "short",
	Long:  "long",
	Run: func(cmd *cobra.Command, args []string) {
		println("HEAD")
	},
}

func init() {
	rootCommand.AddCommand(versionCmd)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
