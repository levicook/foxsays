package router

import (
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
		"home",
		"GET", "/", ping,
	},

	{
		"ping",
		"GET", "/ping", ping,
	},
}
