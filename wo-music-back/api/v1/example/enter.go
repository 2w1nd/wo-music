package example

import (
	"github.com/wo-music-back/server/serivce"
)

type ApiGroup struct {
	AdminApi
}

var adminService = serivce.ServiceGroupApp.ExampleServiceGroup.AdminService
