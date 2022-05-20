package config

import (
	// "github.com/Etwodev/Doctorate/server/helpers"
	"github.com/Etwodev/Doctorate/server/router"
)

func NewRouter(status bool) router.Router {
	return router.NewRouter(initRoutes(), status)
}

func initRoutes() []router.Route {
	return []router.Route{
		router.NewGetRoute("/ABCDEFGHIJKLMNO/config/official/network_config", true, true, NetworkConfigGetRoute),
	}
}