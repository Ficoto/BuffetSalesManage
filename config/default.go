package config

import (
	"gitlab.xinghuolive.com/Backend-Go/camel/db"
	"gitlab.xinghuolive.com/Backend-Go/camel/spark"
)

// base namespace
var (
	ProjectName = "buffet"
	Env         = spark.LoadEnv()
	Schema      = ProjectName
	EnvPrefix   = DevPrefix
	DevPrefix   = "dev"

	RedisDBName   = 0
	RedisConnConf = db.LoadRedisConf()

	MongoDBName   = ProjectName
	MongoConnConf = db.LoadMongoConf()


)