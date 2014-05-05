package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testConfig = `

[httpd]
addr = ":4080"

[mongo]
dial = "127.0.0.1"
database = "foxsays"
`

func Test_load(t *testing.T) {
	assert.Equal(t, Httpd.Addr, "")
	assert.Equal(t, Repos.Dial, "")
	assert.Equal(t, Repos.Database, "")

	load(strings.NewReader(testConfig))

	assert.Equal(t, Httpd.Addr, ":4080")
	assert.Equal(t, Repos.Dial, "127.0.0.1")
	assert.Equal(t, Repos.Database, "foxsays")
}
