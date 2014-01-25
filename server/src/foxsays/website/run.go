package website

import (
	"foxsays/config"
	"foxsays/log"
	"foxsays/pages"
	"foxsays/website/router"
	"github.com/spf13/cobra"
	"net/http"
)

func Run(_ *cobra.Command, _ []string) {

	config.Load()
	log.FatalIf(pages.LoadTemplates(config.Website.Assets))

	log.Printf("website: listening at %v", config.Website.HttpAddr)
	log.FatalIf(http.ListenAndServe(config.Website.HttpAddr, router.New()))
}
