package config

var Website website

type website struct {
	Assets   string
	HttpAddr string `toml:"http_addr"`
}
