package router

import (
	"foxsays/website/pages/public_home"
	"net/http"
)

type route struct {
	name    string
	method  string
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

var routes = []route{

	{
		"public_home",
		"GET", "/", public_home.Page,
	},

	{
		"ping",
		"GET", "/ping", ping,
	},
}
