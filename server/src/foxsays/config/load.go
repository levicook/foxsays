package config

import (
	"foxsays/log"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

func Load() {
	file, err := os.Open(File)
	log.PanicIf(err)
	log.PanicIf(load(file))
}

func load(r io.Reader) (err error) {
	_, err = toml.DecodeReader(r, &struct {
		Httpd    *httpdSection    `toml:"httpd"`
		Mongo    *mongoSection    `toml:"mongo"`
		Password *passwordSection `toml:"password"`
	}{
		Httpd:    &Httpd,
		Mongo:    &mongo,
		Password: &pword,
	})

	return
}
