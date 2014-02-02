package config

import (
	"github.com/levicook/go-detect"
	"os"
)

var (
	AppRoot = detect.String(os.Getenv("APP_ROOT"), "./")
	AppEnv  = detect.String(os.Getenv("APP_ENV"), "development")
	File    string
)
