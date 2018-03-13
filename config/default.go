package config

import (
	"BuffetSalesManage/BuffetSalesManage.git/utils"
)

// base namespace
var (
	ProjectName = "buffet"
	Schema      = ProjectName

	MongoDBName   = ProjectName
	MongoConnConf = utils.LoadMongoConf()
)
