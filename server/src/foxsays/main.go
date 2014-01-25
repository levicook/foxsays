package main

import (
	"fmt"
	"foxsays/config"
	"foxsays/httpd"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var rootCommand = &cobra.Command{
	Use:   "foxsays",
	Short: "foxsays is a twitter clone",
	Long:  ``,
}

func init() {

	rootCommand.PersistentFlags().StringVarP(
		&config.File,
		"config", "c",
		path.Join(config.AppRoot, "config", fmt.Sprintf("%s.toml", config.AppEnv)),
		"")

	rootCommand.AddCommand(&cobra.Command{
		Use:   "httpd",
		Short: "short help on httpd ...",
		Long:  "longer help on httpd ...",
		Run:   httpd.Run,
	})

}

func main() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
