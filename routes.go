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
		"MemInfo",
		"GET",
		"/memory",
		Memory,
	},
	Route{
		"CPUInfo",
		"GET",
		"/cpu",
		CPU,
	},
	Route{
		"Disk",
		"GET",
		"/disks",
		Disk,
	},
	Route{
		"LoadAverage",
		"GET",
		"/load",
		LoadAverage,
	},
	Route{
		"Network",
		"GET",
		"/network",
		Network,
	},
	Route{
		"Host Information",
		"GET",
		"/hostinfo",
		Info,
	},
	Route{
		"Process Information",
		"Get",
		"/processes",
		Processes,
	},
}
