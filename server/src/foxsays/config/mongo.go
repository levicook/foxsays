package config

type mongoSection struct {
	Dial     string `toml:"dial"`
	Database string `toml:"database"`
}
