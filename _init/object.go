package initial

import (
	conf "server/_sys/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	O_SERVER_CONFIG *conf.ServerConfig
	O_LOG           *redis.Client
	O_DB_PG         *gorm.DB
	O_REDIS         *redis.Client
)

// func initConfAndDB() {
// 	LoadServerConfig()
// 	InitGormPg()
// 	InitRedis()
// }
