package utils

import (
	"gitlab.xinghuolive.com/Backend-Go/camel/spark"
	"gitlab.xinghuolive.com/Backend-Go/orca/config"

)

func EnvAbbr(env string) string {
	switch env {
	case spark.Production:
		return config.ProdPrefix
	case spark.Preview:
		return config.PrePrefix
	case spark.Test:
		return config.TestPrefix
	default:
		return config.DevPrefix
	}
}

func GetExDomain(env string) string {
	switch env {
	case spark.Production:
		return "https://www.xiaozhibo.com"
	case spark.Preview:
		return "https://pre.xiaozhibo.com"
	case spark.Test:
		return "https://test.xiaozhibo.com"
	default:
		return "https://dev.xiaozhibo.com"
	}
}
