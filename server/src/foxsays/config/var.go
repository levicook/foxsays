package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/levicook/go-detect"
)

var (
	AppRoot = detect.String(os.Getenv("APP_ROOT"), "./")
	AppEnv  = detect.String(os.Getenv("APP_ENV"), "development")

	Default = path.Join(AppRoot, "config", fmt.Sprintf("%s.toml", AppEnv))
	File    string

	Assets assets
	Repos  repositories

	Httpd httpdSection
	mongo mongoSection
	pword passwordSection
)

var (
	hasSuffix = strings.HasSuffix
)
