package main

import (
	"flag"
	"foxsays/config"
	"foxsays/httpd"
	"foxsays/log"

	"github.com/robmerrell/comandante"
)

func main() {
	foxsays := comandante.New("foxsays", "")
	foxsays.IncludeHelp()

	registerCommand := func(c *comandante.Command) {
		c.FlagInit = wrappedFlagInit(c.Name, c.FlagInit)
		foxsays.RegisterCommand(c)
	}

	registerCommand(httpd.NewCommand())

	log.PanicIf(foxsays.Run())
}

func wrappedFlagInit(name string, innerInit func(fs *flag.FlagSet)) func(fs *flag.FlagSet) {
	return func(fs *flag.FlagSet) {
		log.SetPrefix(name)

		fs.StringVar(&config.File, "config", config.Default, "")
		if innerInit != nil {
			innerInit(fs)
		}

		config.Load()
	}
}
