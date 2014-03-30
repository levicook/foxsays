package website

import (
	"foxsays/config"
	"foxsays/log"
	"foxsays/website/router"
	"net/http"

	"github.com/spf13/cobra"
)

func Run(c *cobra.Command, _ []string) {
	log.SetPrefix(c.Name())

	config.Load()

	config.Mongo.Open()
	defer config.Mongo.Close()

	config.Session.Init()
	config.Website.Init()

	log.Printf("listening at %v", config.Website.HttpAddr)
	log.FatalIf(http.ListenAndServe(config.Website.HttpAddr, router.New()))
}
