package config

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_load(t *testing.T) {

	load(strings.NewReader(`
[website]
http_addr = ":1234"
`))

	assert.Equal(t, Website.HttpAddr, ":1234")
}
