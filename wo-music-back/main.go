package main

import (
	"github.com/wo-music-back/server/core"
	"github.com/wo-music-back/server/global"
	"github.com/wo-music-back/server/initialize"
)

func main()  {
	global.VP = core.Viper()	// 初始化viper
	global.DB = initialize.Gorm()	// gorm连接数据库
	if global.DB != nil {
		initialize.MysqlTables(global.DB) // 初始化表
	}
	core.RunWindowsServer()
}

