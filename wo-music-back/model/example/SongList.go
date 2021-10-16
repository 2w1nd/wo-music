package example

import "github.com/wo-music-back/server/global"

type SongList struct {
	global.MODEL
	SongListNum string `json:"songListNum" form:"songListNum" gorm:"comment:歌单编号"`
	SongListName string `json:"songListName" form:"songListName" gorm:"comment:歌单名称"`
	PosterUrl string `json:"posterUrl" form:"posterUrl" gorm:"comment:封面url"`
	CreaterNum string `json:"createrNum" form:"createrNum" gorm:"comment:创建者编号"`
	Remark string `json:"remark" form:"remark" gorm:"comment:备注"`
}
