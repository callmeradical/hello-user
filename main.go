package main

import (
	"log"

	"github.com/2ndWatch/golibs/router"
	"github.com/2ndWatch/golibs/server"
	"github.com/2ndWatch/golibs/version"
	"github.com/callmeradical/hello-user/api"
)

func main() {
	log.Println("Initializing API...")
	version.Release("2.0.0", "Hello User Service")

	root := router.Routes{
		router.Route{"Health", "GET", "/health", api.Healthz},
		router.Route{"/", "GET", "/", api.Index},
	}

	rtr := router.NewRouter(root, map[string]router.Routes{})

	server.Start("8080", rtr)
}
