package initialize

import (
	"github.com/gin-gonic/gin"
	_router "github.com/wo-music-back/server/router"
)

// Routers
// @Description: 初始化总路由
// @return: *gin.Engine
func Routers() *gin.Engine {
	var Router = gin.Default()

	router := _router.GroupApp.Example
	PrivateGroup := Router.Group("")
	PrivateGroup.Use()
	{
		router.InitAdminRouter(PrivateGroup)
	}
	return Router
}
