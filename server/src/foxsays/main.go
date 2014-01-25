package main

import (
	"fmt"
	"foxsays/config"
	"foxsays/httpd"
	"foxsays/log"
	"github.com/spf13/cobra"
	"path"
)

func main() {
	rootCommand := &cobra.Command{
		Use:   "foxsays",
		Short: "foxsays is a twitter clone"}

	rootCommand.PersistentFlags().StringVarP(
		&config.File, "config", "c",
		path.Join(config.AppRoot, "config", fmt.Sprintf("%s.toml", config.AppEnv)), "")

	rootCommand.AddCommand(&cobra.Command{
		Use: "httpd",
		Run: httpd.Run})

	log.FatalIf(rootCommand.Execute())
}
