package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testConfig = `

[httpd]
addr = ":3080"
session_auth = "im-a-secret"
session_name = "foxsays-test-session"

[mongo]
dial = "127.0.0.1"
database = "foxsays-test-database"

[password]
salt = "im-another-secret"
`

func Test_load(t *testing.T) {
	load(strings.NewReader(testConfig))

	assert.Equal(t, Httpd.Addr, ":3080")
	assert.Equal(t, Httpd.SessionAuth, "im-a-secret")
	assert.Equal(t, Httpd.SessionName, "foxsays-test-session")

	assert.Equal(t, pword.Salt, "im-another-secret")

	assert.Equal(t, mongo.Dial, "127.0.0.1")
	assert.Equal(t, mongo.Database, "foxsays-test-database")
}
