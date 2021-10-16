package example

import "github.com/wo-music-back/server/global"

type Collection struct {
	global.MODEL
	UserId int `json:"userId" form:"userId" gorm:"comment:用户id"`
	SongId int `json:"songId" form:"SongId" gorm:"comment:歌曲id"`
}
