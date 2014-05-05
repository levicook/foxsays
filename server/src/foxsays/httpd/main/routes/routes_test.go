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
		{"GET", "/dashboard", "main_dashboard", vars{}},

		{"GET", "/images/asdf", "main_show_image", vars{
			"imageId": "asdf",
		}},

		{"GET", "/images/asdf/meta", "main_show_image_meta", vars{
			"imageId": "asdf",
		}},

		{"GET", "/images/asdf/download", "main_show_image_download", vars{
			"imageId": "asdf",
		}},

		{"GET", "/api/users/10", "main_api_show_user", vars{
			"userId": "10",
		}},

		{"GET", "/api/users/me", "main_api_show_current_user", vars{}},

		{"POST", "/api/users/signin", "main_api_signin_user", vars{}},

		{"POST", "/api/users/signup", "main_api_signup_user", vars{}},
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
