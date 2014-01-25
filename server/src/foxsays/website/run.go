package website

import (
	"foxsays/config"
	"foxsays/website/router"
	"foxsays/log"
	"github.com/spf13/cobra"
	"net/http"
)

func Run(_ *cobra.Command, _ []string) {
	config.Load()
	log.Printf("website: listening at %v", config.Website.HttpAddr)
	log.FatalIf(http.ListenAndServe(config.Website.HttpAddr, router.New()))
}
