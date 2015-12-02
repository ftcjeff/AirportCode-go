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
		"ShowById",
		"GET",
		"/airport/id/{Id}",
		ShowById,
	},

	Route{
		"City",
		"GET",
		"/airport/city/{city}",
		City,
	},

	Route{
		"Country",
		"GET",
		"/airport/country/{country}",
		Country,
	},

	Route{
		"Country",
		"GET",
		"/airport/country/{country}",
		Country,
	},
}
