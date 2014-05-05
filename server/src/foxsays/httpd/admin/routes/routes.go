package routes

import (
	imagesAPI "foxsays/httpd/admin/api/images"

	"foxsays/httpd/route"
)

func Routes() route.Routes {
	return route.Routes{

		{ // imagesAPI
			"admin_api_create_image",
			"POST", "/admin/api/images", imagesAPI.Create,
		},
	}
}
