package main

import "net/http"

// match url
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
		addStudent,
	},
	Route{
		"AddClass",
		"POST",
		"/register-class",
		addClass,
	},
	Route{
		"GetSumScore",
		"GET",
		"/get-class-total-score/{id}",
		getTotalScore,
	},
	Route{
		"GetTeacherFromScore",
		"GET",
		"/get-top-teacher",
		getTeacherFromScore,
	},
}
