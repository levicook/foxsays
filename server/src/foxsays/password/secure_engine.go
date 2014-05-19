package password

import (
	"crypto/sha256"
	"crypto/subtle"

	"code.google.com/p/go.crypto/pbkdf2"
)

func NewSecureEngine(salt []byte) Engine {
	return secureEngine{salt: salt}
}

type secureEngine struct{ salt []byte }

func (e secureEngine) Digest(pass string) []byte {
	return e.digest([]byte(pass))
}

func (e secureEngine) Equal(pass1, pass2 []byte) bool {
	return subtle.ConstantTimeCompare(pass1, pass2) == 1
}

func (e secureEngine) digest(pass []byte) []byte {
	return pbkdf2.Key(pass, e.salt, 4096, sha256.Size, sha256.New)
}
