package example

import "github.com/wo-music-back/server/global"

type SongCategory struct {
	global.MODEL
	ParentId int `json:"parentId" form:"parentId" gorm:"comment:父级ID"`
	Sort int `json:"sort" form:"sort" gorm:"comment:排序"`
	Remark string `json:"remark" form:"remark" gorm:"comment:备注"`
}
