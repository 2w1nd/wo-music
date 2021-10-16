package example

import "github.com/wo-music-back/server/global"

type SongSongListRela struct {
	global.MODEL
	SongListId int `json:"songListId" form:"songListId" gorm:"comment:歌单id"`
	SongId int `json:"songId" form:"songId" gorm:"comment:歌曲id"`
}
