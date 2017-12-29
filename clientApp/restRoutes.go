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
		"GetIDs",
		"GET",
		"/ids",
		GetIDs,
	},
	Route{
		"GetNewID",
		"GET",
		"/newid",
		GetNewID,
	},
	Route{
		"GetBlindAndSendToSign",
		"GET",
		"/blindandsendtosign/{keyid}",
		GetBlindAndSendToSign,
	},
	Route{
		"GetVerify",
		"GET",
		"/verify/{keyid}",
		GetVerify,
	},
	Route{
		"GetDelete",
		"GET",
		"/delete/{keyid}",
		GetDelete,
	},
}
