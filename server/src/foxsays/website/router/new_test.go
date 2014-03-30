package router

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	testCases := []struct {
		method, path, name string
	}{
		{"GET", "/", "public_home"},
		{"GET", "/dashboard", "dashboard"},
		{"GET", "/forgot_password", "forgot_password"},
		{"GET", "/ping", "ping"},
		{"GET", "/settings", "settings"},
		{"GET", "/sign_in", "sign_in"},
		{"GET", "/sign_out", "sign_out"},
	}

	router := New()
	assert.NotNil(t, router)

	for _, testCase := range testCases {
		req, err := http.NewRequest(testCase.method, testCase.path, nil)
		assert.Nil(t, err)

		var match mux.RouteMatch
		if assert.True(t, router.Match(req, &match), "expected a route") {
			assert.Equal(t, testCase.name, match.Route.GetName())
		}
	}
}
