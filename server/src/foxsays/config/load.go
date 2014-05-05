package config

import (
	"io"
	"os"
	"foxsays/log"

	"github.com/BurntSushi/toml"
)

func Load() {
	file, err := os.Open(File)
	log.PanicIf(err)
	log.PanicIf(load(file))
}

func load(r io.Reader) (err error) {
	_, err = toml.DecodeReader(r, &struct {
		Httpd *httpd `toml:"httpd"`
		Mongo *mongo `toml:"mongo"`
	}{
		Httpd: &Httpd,
		Mongo: &Repos,
	})

	return
}
