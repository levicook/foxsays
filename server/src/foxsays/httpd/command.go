package httpd

import (
	"net/http"
	"foxsays/config"
	"foxsays/httpd/router"
	"foxsays/httpd/utils/gzip"
	"foxsays/log"

	"github.com/robmerrell/comandante"
)

func NewCommand() *comandante.Command {
	return comandante.NewCommand("httpd", "run http server", func() error {

		assets := config.Assets.Load()
		http.Handle(assets.Prefix(), gzip.Handler(assets.Handler()))

		http.Handle("/", gzip.Handler(router.New()))

		log.FatalIf(config.Repos.Open())
		defer config.Repos.Close()

		log.Printf("listening at %v", config.Httpd.Addr)
		return http.ListenAndServe(config.Httpd.Addr, nil)
	})
}
