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
		"/",
		Index,
	},
	Route{
		"AddStudent",
		"POST",
		"/register-student",
		AddStudent,
	},
	Route{
		"AddClass",
		"POST",
		"/register-class",
		AddClass,
	},
	Route{
		"GetSumScore",
		"GET",
		"/get-class-total-score/{id}",
		GetSumScore,
	},
	Route{
		"GetTeacherFromScore",
		"GET",
		"/get-top-teacher",
		GetTeacherFromScore,
	},
}
