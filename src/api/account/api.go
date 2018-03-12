package account

import (
	"BuffetSalesManage/BuffetSalesManage.git/router"
	"net/http"
)

var ExRouter = router.ModuleRouter{
	URLPrefix: "/api/account",
	SubRouters: []router.SubRouter{
		{
			Name:        "register",
			Methods:     []string{http.MethodGet},
			Pattern:     "/register",
			HandlerFunc: register,
		},
	},
}

type registerResponse struct {
	IsSuccess bool `json:"is_success"`
}

func register(w http.ResponseWriter, r *http.Request) {
	var response registerResponse
	response.IsSuccess = true
	router.JSONResp(w, http.StatusOK, response)
}
