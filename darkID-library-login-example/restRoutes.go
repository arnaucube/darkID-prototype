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
		"GetProof",
		"POST",
		"/getproof",
		GetProof,
	},
	Route{
		"AnswerProof",
		"POST",
		"/answerproof",
		AnswerProof,
	},
}
