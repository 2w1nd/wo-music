package example

import "github.com/wo-music-back/server/global"

type Comment struct {
	global.MODEL
	UserId int `json:"userId" form:"userId" gorm:"comment:用户id"`
	SongId int `json:"songId" form:"songId" gorm:"comment:歌曲id"`
	Content string `json:"content" form:"content" gorm:"comment:评论内容"`
}
