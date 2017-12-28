package main

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetServer",
		"GET",
		"/getserver",
		GetServer,
	},
	Route{
		"IDs",
		"GET",
		"/ids",
		IDs,
	},
	Route{
		"NewID",
		"GET",
		"/newid",
		NewID,
	},
	Route{
		"BlindAndSendToSign",
		"GET",
		"/blindandsendtosign/{pubK}",
		BlindAndSendToSign,
	},
	Route{
		"Verify",
		"GET",
		"/verify/{pubK}",
		Verify,
	},
}
