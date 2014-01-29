package router

import (
	"foxsays/github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ping(t *testing.T) {
	w := httptest.NewRecorder()
	r, e := http.NewRequest("GET", "/ping", nil)
	assert.Nil(t, e)

	ping(w, r)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, pong, w.Body.Bytes())
}
