package base

import (
	"log"
	"sync"

	"gitlab.xinghuolive.com/Backend-Go/camel/spark"
	"gitlab.xinghuolive.com/Backend-Go/orca/config"
)

var (
	icOnce sync.Once
)

// InitialConfigParam - Initialize config params when program start
func InitialConfigParam() {
	icOnce.Do(func() {
		// 初始化环境变量前缀
		switch config.Env {
		case spark.Production:
			config.EnvPrefix = config.ProdPrefix
		case spark.Preview:
			config.EnvPrefix = config.PrePrefix
			config.MongoDBName = "pre-kangaroo"
		case spark.Test:
			config.EnvPrefix = config.TestPrefix
		default:
			config.EnvPrefix = config.DevPrefix
		}
		log.Printf("INFO(%s): Echo sys environment!", config.Env)
	})
}
