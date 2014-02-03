package config

import (
	"foxsays/pages"
	"path"
)

var Website website

type website struct {
	Assets   string
	HttpAddr string `toml:"http_addr"`

	Pages pages.PageSet `toml:"-"`
}

func (w *website) Init() {
	w.Pages = pages.LoadPages(path.Join(w.Assets, `website/pages`))
}

func (w website) GetPage(name string) pages.Page {
	return w.Pages.Get(name)
}
