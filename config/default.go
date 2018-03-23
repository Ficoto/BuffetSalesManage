package config

import (
	"BuffetSalesManage/BuffetSalesManage/utils"
)

// base namespace
var (
	ProjectName = "buffet"
	Schema      = ProjectName

	MongoDBName   = ProjectName
	MongoConnConf = utils.LoadMongoConf()
)
