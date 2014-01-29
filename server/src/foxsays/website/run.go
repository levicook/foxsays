package website

import (
	"foxsays/config"
	"foxsays/log"
	"foxsays/pages"
	"foxsays/website/router"
	"foxsays/github.com/spf13/cobra"
	"net/http"
	"path"
)

func Run(_ *cobra.Command, _ []string) {
	config.Load()

	templatePath := path.Join(config.Website.Assets, `website/pages`)
	log.FatalIf(pages.LoadTemplates(templatePath))

	log.Printf("website: listening at %v", config.Website.HttpAddr)
	log.FatalIf(http.ListenAndServe(config.Website.HttpAddr, router.New()))
}
