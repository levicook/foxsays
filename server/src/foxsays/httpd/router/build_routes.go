package router

import (
	admin "foxsays/httpd/admin/routes"
	main "foxsays/httpd/main/routes"

	"foxsays/httpd/route"
)

func buildRoutes() (routes route.Routes) {
	routes = append(routes, admin.Routes()...)
	routes = append(routes, main.Routes()...)
	return
}
