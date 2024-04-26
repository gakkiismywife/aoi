package utils

import (
	"AiServer/configs"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

var Rdb *redis.Client

var rdbOnce sync.Once

func InitRedisClient(c *configs.Redis) {
	rdbOnce.Do(func() {
		Rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
			Password: c.Password, // no password set
			DB:       c.Index,    // use default DB
		})
	})
}
