package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ping(t *testing.T) {
	w := httptest.NewRecorder()
	r, e := http.NewRequest("GET", "/ping", nil)
	assert.Nil(t, e)

	ping(w, r)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, pong, w.Body.Bytes())
}
