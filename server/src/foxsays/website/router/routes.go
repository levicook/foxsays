package router

import (
	"foxsays/website/filters"
	"foxsays/website/pages/dashboard"
	"foxsays/website/pages/forgot_password"
	"foxsays/website/pages/public_home"
	"foxsays/website/pages/settings"
	"foxsays/website/pages/sign_in"
)

var routes = []route{
	{
		"home",
		"GET", "/", filterSet{
			// todo: rename public_home -> signed_out_home
			// todo: rename dashboard   -> signed_in_home
			public_home.Page,
			dashboard.Page,
		},
	},
	{
		"forgot_password",
		"GET", "/forgot_password", filterSet{
			filters.EnsureSignedOut,
			forgot_password.Page,
		},
	},
	{
		"ping",
		"GET", "/ping", filterSet{ping},
	},
	{
		"settings",
		"GET", "/settings", filterSet{
			filters.EnsureSignedIn,
			settings.Page,
		},
	},
	{
		"sign_in",
		"GET", "/sign_in", filterSet{
			filters.EnsureSignedOut,
			sign_in.Page,
		},
	},
	{
		"sign_out",
		"GET", "/sign_out", filterSet{
			filters.SignOut,
		},
	},
}
