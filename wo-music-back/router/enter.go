package router

import "github.com/wo-music-back/server/router/example"

type AllRouterGroup struct {
	Example example.RouterGroup
}

var GroupApp = new(AllRouterGroup)
