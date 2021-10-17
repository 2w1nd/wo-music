package serivce

import "github.com/wo-music-back/server/serivce/example"

type AllServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(AllServiceGroup)
