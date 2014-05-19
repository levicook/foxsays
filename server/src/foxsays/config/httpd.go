package config

import "github.com/gorilla/sessions"

type httpdSection struct {
	Addr        string `toml:"addr"`
	SessionAuth string `toml:"session_auth"`
	SessionName string `toml:"session_name"`
}

func (h httpdSection) NewSessionStore() sessions.Store {
	return sessions.NewCookieStore([]byte(h.SessionAuth))
}
