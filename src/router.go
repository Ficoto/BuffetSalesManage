package src

import (
	"github.com/gorilla/mux"
	"BuffetSalesManage/BuffetSalesManage.git/src/api/account"
	"BuffetSalesManage/BuffetSalesManage.git/router"
)

// KangarooRouter -
var BuffetSalesRouter = router.BaseRouter{
	R: mux.NewRouter(),
	ModuleRouters: []router.ModuleRouter{
		account.ExRouter,
	},
}
