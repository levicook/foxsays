package routes

import (
	imagesAPI "foxsays/httpd/admin/api/images"

	"foxsays/httpd/admin/pages/dashboard"
	"foxsays/httpd/admin/pages/home"
	"foxsays/httpd/route"
	"foxsays/httpd/sessions"
)

func Routes() route.Routes {
	return route.Routes{

		{ // home, dashboard, logout
			"admin_home",
			"GET", "/admin", home.Show,
		}, {
			"admin_dashboard",
			"GET", "/admin/dashboard", dashboard.Show,
		}, {
			"admin_logout",
			"GET", "/admin/logout", sessions.Logout("/admin"),
		},

		{ // imagesAPI
			"admin_api_create_image",
			"POST", "/admin/api/images", imagesAPI.Create,
		},
	}
}
