package account

import "BuffetSalesManage/router"

var ExRouter = router.ModuleRouter{
	URLPrefix: "api/account",
	SubRouters:[]router.SubRouter{
		{
			Name:"reg"
		},
	},
}