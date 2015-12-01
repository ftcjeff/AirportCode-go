package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/airport",
		Index,
	},

	Route{
		"Show",
		"GET",
		"/airport/{airportCode}",
		Show,
	},

	Route{
		"City",
		"GET",
		"/airport/city/{city}",
		City,
	},

	Route{
		"State",
		"GET",
		"/airport/state/{state}",
		State,
	},

	Route{
		"Country",
		"GET",
		"/airport/country/{country}",
		Country,
	},

	Route{
		"Remove",
		"DELETE",
		"/airport/{airportCode}",
		Remove,
	},

	Route{
		"Create",
		"POST",
		"/airport",
		Create,
	},
}
