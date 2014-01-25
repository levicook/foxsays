package config

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_load(t *testing.T) {

	load(strings.NewReader(`
[httpd]
http_addr = ":1234"
`))

	assert.Equal(t, Httpd.HttpAddr, ":1234")
}
