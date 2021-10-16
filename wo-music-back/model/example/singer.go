package example

import "github.com/wo-music-back/server/global"

type Single struct {
	global.MODEL
	SingerNum string `json:"singerNum" form:"singerNum" gorm:"comment:歌手编号"`
	SingerName string `json:"singerName" form:"singerName" gorm:"comment:歌手名称"`
	Sex string `json:"sex" form:"sex" gorm:"comment:性别"`
	Birthday string `json:"birthday" form:"birthday" gorm:"comment:生日"`
	From string `json:"from" form:"from" gorm:"comment:来自"`
	Info string `json:"info" form:"info" gorm:"comment:介绍"`
}
