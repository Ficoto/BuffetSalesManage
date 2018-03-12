package config

import (
	"gitlab.xinghuolive.com/Backend-Go/camel/db"
)

// base namespace
var (
	ProjectName = "buffet"
	Schema      = ProjectName
	EnvPrefix   = DevPrefix
	DevPrefix   = "dev"

	RedisDBName   = 0
	RedisConnConf = db.LoadRedisConf()

	MongoDBName   = ProjectName
	MongoConnConf = db.LoadMongoConf()
)
