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
		"root",
		"GET", "/", ping,
	},

	{
		"ping",
		"GET", "/ping", ping,
	},
}
