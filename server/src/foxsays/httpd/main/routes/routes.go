package routes

import (
	usersAPI "foxsays/httpd/main/api/users"

	"foxsays/httpd/main/images"
	"foxsays/httpd/main/pages/dashboard"
	"foxsays/httpd/main/pages/home"
	"foxsays/httpd/route"
)

func Routes() route.Routes {
	return route.Routes{

		// dashboard
		{
			"main_dashboard",
			"GET", "/dashboard", dashboard.Show,
		},

		// home
		{
			"main_home",
			"GET", "/", home.Show,
		},

		// images
		{
			"main_show_image",
			"GET", "/images/{imageId}", images.Show,
		}, {
			"main_show_image_download",
			"GET", "/images/{imageId}/download", images.Download,
		}, {
			"main_show_image_meta",
			"GET", "/images/{imageId}/meta", images.Meta,
		},

		// usersAPI
		{
			"main_api_signup_user",
			"POST", "/api/users/signup", usersAPI.Signup,
		}, {
			"main_api_signin_user",
			"POST", "/api/users/signin", usersAPI.Signin,
		}, {
			"main_api_show_current_user",
			"GET", "/api/users/me", usersAPI.ShowCurrent,
		}, {
			"main_api_show_user",
			"GET", "/api/users/{userId:[0-9]+}", usersAPI.Show,
		},
	}
}
