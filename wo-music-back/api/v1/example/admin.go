package example

import (
	"github.com/gin-gonic/gin"
	"github.com/wo-music-back/server/model/common/response"
	"github.com/wo-music-back/server/model/example"
	_ "github.com/wo-music-back/server/serivce/example"
	"log"
)

type AdminApi struct {
}

// CreateAdmin
// @Description: 创建管理员
// @receiver: e
// @param: c
func (e *AdminApi) CreateAdmin(c *gin.Context) {
	var admin example.Admin
	_ = c.ShouldBindJSON(&admin)
	//TODO 校验+jwt认证
	log.Println(admin.AdminPassword)
	if err := adminService.CreateAdmin(admin); err != nil {
		log.Println("创建失败=====", c)
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAdmin
// @Description: 删除管理员
// @receiver: e
// @param: c
func (e *AdminApi) DeleteAdmin(c *gin.Context) {
	var admin example.Admin
	_ = c.ShouldBindJSON(&admin)
	if err := adminService.DeleteAdmin(admin); err != nil {
		log.Println("删除失败=====", c)
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateAdmin
// @Description: 更新管理员（全部字段都要更新，不传的默认为空）
// @receiver: e
// @param: c
func (e *AdminApi) UpdateAdmin(c *gin.Context) {
	var admin example.Admin
	_ = c.ShouldBindJSON(&admin)
	if err := adminService.UpdateAdmin(&admin); err != nil {
		log.Println("更新失败=====", c)
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
