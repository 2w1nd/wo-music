package example

import "github.com/wo-music-back/server/global"

type Song struct {
	global.MODEL
	SongNum string `json:"songNum" form:"songNum" gorm:"comment:歌曲编号"`
	SongName string `json:"songName" form:"songName" gorm:"comment:歌曲名称"`
	SingerNum int `json:"singerNum" form:"singerNum" gorm:"comment:歌手编号"`
	SongUrl string `json:"songUrl" form:"songUrl" gorm:"comment:歌曲url"`
	PosterUrl string `json:"posterUrl" form:"posterUrl" gorm:"comment:封面url"`
	OneCategoryId int `json:"oneCategoryId" form:"oneCategoryId" gorm:"comment:一级分类id"`
	TwoCategoryId int `json:"TwoCategoryId" form:"TwoCategoryId" gorm:"comment:二级分类id"`
	ThreeCategoryId int `json:"ThreeCategoryId" form:"ThreeCategoryId" gorm:"comment:三级分类id"`
	Status int `json:"status" form:"status" gorm:"comment:状态"`
}
