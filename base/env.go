package base

import (
	"log"
	"sync"
	"BuffetSalesManage/BuffetSalesManage.git/config"
)

var (
	icOnce sync.Once
)

// InitialConfigParam - Initialize config params when program start
func InitialConfigParam() {
	icOnce.Do(func() {
		// 初始化环境变量前缀
		switch config.Env {
		default:
			config.EnvPrefix = config.DevPrefix
		}
		log.Printf("INFO(%s): Echo sys environment!", config.Env)
	})
}
