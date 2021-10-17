package v1

import "github.com/wo-music-back/server/api/v1/example"

type AllApiGroup struct {
	ApiGroup example.ApiGroup
}

var ApiGroupApp = new(AllApiGroup)