package httpd

import (
	"foxsays/config"
	"foxsays/httpd/router"
	"foxsays/log"
	"github.com/spf13/cobra"
	"net/http"
)

func Run(_ *cobra.Command, _ []string) {

	config.Load()

	http.Handle("/", router.New())
	log.Printf("httpd: listening at %v", config.Httpd.HttpAddr)
	log.FatalIf(http.ListenAndServe(config.Httpd.HttpAddr, nil))

}
