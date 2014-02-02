package config

import (
	"foxsays/log"
	"github.com/BurntSushi/toml"
	"io"
	"os"
)

func Load() {
	f, err := os.Open(File)
	log.FatalIf(err)
	load(f)
}

func load(r io.Reader) {
	_, err := toml.DecodeReader(r, &struct {
		Session *session
		Website *website
	}{
		Session: &Session,
		Website: &Website,
	})

	log.FatalIf(err)
}
