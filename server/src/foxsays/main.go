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
	rootCmd := &cobra.Command{
		Use:   "foxsays",
		Short: "foxsays is a twitter clone",
	}

	rootCmd.PersistentFlags().StringVarP(
		&config.File, "config", "c",
		path.Join(config.AppRoot, "config", fmt.Sprintf("%s.toml", config.AppEnv)),
		"")

	rootCmd.AddCommand(&cobra.Command{
		Use: "httpd",
		Run: httpd.Run,
	})

	log.FatalIf(rootCmd.Execute())
}
