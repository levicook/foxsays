package config

import (
	"foxsays/github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_load(t *testing.T) {

	load(strings.NewReader(`
[website]
assets = "/opt/foxsays/assets"
http_addr = ":1234"
`))

	assert.Equal(t, Website.Assets, "/opt/foxsays/assets")
	assert.Equal(t, Website.HttpAddr, ":1234")
}
