package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_load(t *testing.T) {

	load(strings.NewReader(`
[mongo]
database = "foxsays"
dial = "127.0.0.1"

[session]
new_authentication_key = "new_authentication_key"
new_encryption_key     = "new_encryption_key"
old_authentication_key = "old_authentication_key"
old_encryption_key     = "old_encryption_key"

[website]
assets = "/opt/foxsays/assets"
http_addr = ":1234"
`))
	assert.Equal(t, Mongo.Database, "foxsays")
	assert.Equal(t, Mongo.Dial, "127.0.0.1")

	assert.Equal(t, Session.NewAuthenticationKey, "new_authentication_key")
	assert.Equal(t, Session.NewEncryptionKey, "new_encryption_key")
	assert.Equal(t, Session.OldAuthenticationKey, "old_authentication_key")
	assert.Equal(t, Session.OldEncryptionKey, "old_encryption_key")

	assert.Equal(t, Website.Assets, "/opt/foxsays/assets")
	assert.Equal(t, Website.HttpAddr, ":1234")
}
