package config

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Session session

type session struct {
	NewAuthenticationKey string `toml:"new_authentication_key"`
	OldAuthenticationKey string `toml:"old_authentication_key"`
	NewEncryptionKey     string `toml:"new_encryption_key"`
	OldEncryptionKey     string `toml:"old_encryption_key"`

	store sessions.Store `toml:"-"`
}

func (s *session) Init() {
	s.store = sessions.NewCookieStore(
		[]byte(s.NewAuthenticationKey),
		[]byte(s.NewEncryptionKey),
		[]byte(s.OldAuthenticationKey),
		[]byte(s.OldEncryptionKey),
	)
}

func (s session) ForWebsite(r *http.Request) *sessions.Session {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	gs, _ := s.store.Get(r, "website")
	return gs
}
