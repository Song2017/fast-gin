package initial

import (
	"context"
	"fmt"
	conf "server/_sys/model"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisCfg := LoadServerConfig().Redis
	return InitRedisByConfig(redisCfg)
}
func InitRedisByConfig(redisCfg conf.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect ping failed, err:", err)
		panic(err)
	} else {
		fmt.Println("redis connect ping response:", pong)
	}
	return client
}
