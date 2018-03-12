package redis_x

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	"BuffetSalesManage/BuffetSalesManage.git/config"
)

// RedisSentinelOps - sentinel connect info
type RedisSentinelOps struct {
	DB            int
	MasterName    string
	Password      string
	SentinelAddrs []string
}

var pool *redis.Pool
var syncOnce sync.Once

// GetRedisPool - connect to redis pool
func GetRedisPool() *redis.Pool {
	syncOnce.Do(
		func() {
			pool = &redis.Pool{
				MaxIdle:     3,
				MaxActive:   64,
				Wait:        true,
				IdleTimeout: 240 * time.Second,
				Dial: func() (redis.Conn, error) {
					c, err := redis.Dial("tcp", config.RedisConnConf.Hosts[0])
					if err != nil {
						return nil, err
					}
					if _, err := c.Do("AUTH", config.RedisConnConf.Password); err != nil {
						c.Close()
						return nil, err
					}
					if _, err := c.Do("SELECT", config.RedisDBName); err != nil {
						c.Close()
						return nil, err
					}
					return c, nil
				},
			}
		},
	)
	return pool
}
