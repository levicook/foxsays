package config

import (
	"foxsays/log"
	"github.com/BurntSushi/toml"
	"github.com/levicook/go-detect"
	"io"
	"os"
)

var (
	AppRoot = detect.String(os.Getenv("APP_ROOT"), "./")
	AppEnv  = detect.String(os.Getenv("APP_ENV"), "development")

	File string

	Website website
)

type website struct {
	Assets   string
	HttpAddr string `toml:"http_addr"`
}

func Load() {
	f, err := os.Open(File)
	log.FatalIf(err)
	load(f)
}

func load(r io.Reader) {
	_, err := toml.DecodeReader(r, &struct {
		Website *website
	}{
		Website: &Website,
	})

	log.FatalIf(err)
}
