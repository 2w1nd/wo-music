package example

import (
	"github.com/gin-gonic/gin"
	"github.com/wo-music-back/server/api/v1"
)

type AdminRouter struct {
}

func (e *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouter := Router.Group("admin")
	var adminApi = v1.ApiGroupApp.ApiGroup.AdminApi
	{
		adminRouter.POST("admin", adminApi.CreateAdmin) // 创建管理员
	}
}
