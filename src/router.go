package src

import (
	"github.com/gorilla/mux"
	"BuffetSalesManage/BuffetSalesManage/router"
	"BuffetSalesManage/BuffetSalesManage/src/api/businesses.account"
)

// KangarooRouter -
var BuffetSalesRouter = router.BaseRouter{
	R: mux.NewRouter(),
	ModuleRouters: []router.ModuleRouter{
		businesses_account.ExRouter,
	},
}
