package config

import (
	"gitlab.xinghuolive.com/Backend-Go/camel/db"
	"gitlab.xinghuolive.com/Backend-Go/camel/spark"
)

// base namespace
var (
	ProjectName = "orca"
	Env         = spark.LoadEnv()
	Schema      = ProjectName
	EnvPrefix   = DevPrefix
	DevPrefix   = "dev"
	TestPrefix  = "test"
	PrePrefix   = "pre"
	ProdPrefix  = "prod"

	RedisDBName   = 0
	RedisConnConf = db.LoadRedisConf()

	MongoDBName   = ProjectName
	MongoConnConf = db.LoadMongoConf()


)

const (
	DefaultNotLimitSize = 0
	MaxImageDefaultSize = DefaultNotLimitSize
	MaxAudioDefaultSize = DefaultNotLimitSize
	MaxVideoDefaultSize = DefaultNotLimitSize
	MaxOtherDefaultSize = DefaultNotLimitSize
)
