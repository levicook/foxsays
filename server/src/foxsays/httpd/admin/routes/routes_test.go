package routes

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_Routes(t *testing.T) {
	type vars map[string]string

	testCases := []struct {
		method, path, name string
		vars               vars
	}{

		{"POST", "/admin/api/images", "admin_api_create_image", vars{}},
	}

	router := Routes().Router()
	assert.NotNil(t, router)

	for _, testCase := range testCases {
		req, err := http.NewRequest(testCase.method, testCase.path, nil)
		assert.Nil(t, err)

		var match mux.RouteMatch
		if assert.True(t, router.Match(req, &match), "expected a route") {
			if assert.Equal(t, testCase.name, match.Route.GetName()) {
				vars := map[string]string(testCase.vars)
				assert.Equal(t, vars, match.Vars)
			}
		}
	}
}
