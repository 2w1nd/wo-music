package core

import (
	"fmt"
	"github.com/wo-music-back/server/global"
	"github.com/wo-music-back/server/initialize"
	"log"
)

type server interface {
	ListenAndServer() error
}

func RunWindowsServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)

	s := initServer(address, Router)
	fmt.Println(`
		wo-music项目成功启动
		====================================================================
	`)
	log.Println(s.ListenAndServe().Error())
}
