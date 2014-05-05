package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"foxsays/log"
)

type assets struct {
	manifest revManifest
}

type revManifest map[string]string

func (a assets) Get(keys ...string) []string {
	vals := make([]string, len(keys), len(keys))

	for i := range keys {

		v, ok := a.manifest[keys[i]]
		if !ok {
			log.Panicf("asset not found: %q", keys[i])
		}
		vals[i] = path.Join(a.Prefix(), v)

	}

	return vals
}

func (a assets) Handler() http.Handler {
	d := http.Dir(a.path())
	f := http.FileServer(d)
	p := http.StripPrefix(a.Prefix(), f)
	return expires(p)
}

func (a *assets) Load() *assets {
	revManifestPath := path.Join(a.path(), "rev-manifest.json")
	data, err := ioutil.ReadFile(revManifestPath)
	log.PanicIf(err)
	log.PanicIf(json.Unmarshal(data, &a.manifest))
	return a
}

func (a assets) Prefix() string {
	return "/assets/"
}

func (a assets) path() string {
	switch AppEnv {
	case "development":
		return path.Join(AppRoot, "client/dist/assets")
	default:
		return path.Join(AppRoot, "assets")
	}
}
