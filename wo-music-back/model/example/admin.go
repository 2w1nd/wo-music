package example

import "github.com/wo-music-back/server/global"

type Admin struct {
	global.MODEL
	AdminName string `json:"adminName" form:"adminName" gorm:"comment:管理员名"`
	AdminPassword string `json:"adminPassword" form:"adminPassword" gorm:"comment:管理员密码"`
	Type int	`json:"type" from:"type" gorm:"comment:类型"`
}


