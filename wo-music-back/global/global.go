package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/wo-music-back/server/config"
	"gorm.io/gorm"
)

var (
	DB	*gorm.DB
	REDIS	*redis.Client
	CONFIG	config.Server
	VP	*viper.Viper
)
