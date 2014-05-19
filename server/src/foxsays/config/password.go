package config

import "foxsays/password"

type passwordSection struct {
	Salt string `toml:"salt"`
}

func (pws passwordSection) engine() password.Engine {
	return password.NewSecureEngine([]byte(pws.Salt))
}
