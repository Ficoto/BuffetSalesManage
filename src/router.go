package src

import (
	"github.com/gorilla/mux"
	"BuffetSalesManage/BuffetSalesManage/router"
	"BuffetSalesManage/BuffetSalesManage/src/api/businesses.account"
	"BuffetSalesManage/BuffetSalesManage/src/api/consumer.account"
	"BuffetSalesManage/BuffetSalesManage/src/api/commodity"
)

// KangarooRouter -
var BuffetSalesRouter = router.BaseRouter{
	R: mux.NewRouter(),
	ModuleRouters: []router.ModuleRouter{
		businesses_account.ExRouter,
		consumer_account.ExRouter,
		commodity.ExRouter,
	},
}
