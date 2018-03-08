package src

import (
	"github.com/gorilla/mux"
	"gitlab.xinghuolive.com/Backend-Go/orca/router"
)

// KangarooRouter -
var OrcaRouter = router.BaseRouter{
	R: mux.NewRouter(),
	ModuleRouters: []router.ModuleRouter{

	},
}
