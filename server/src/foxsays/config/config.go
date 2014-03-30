package config

import (
	"os"

	"github.com/levicook/go-detect"
)

var (
	AppRoot = detect.String(os.Getenv("APP_ROOT"), "./")
	AppEnv  = detect.String(os.Getenv("APP_ENV"), "development")
	File    string
)
