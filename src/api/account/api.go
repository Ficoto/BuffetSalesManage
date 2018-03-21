package account

import (
	"BuffetSalesManage/BuffetSalesManage.git/router"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"BuffetSalesManage/BuffetSalesManage.git/base/error.code"
	"BuffetSalesManage/BuffetSalesManage.git/model/mongo"
)

var ExRouter = router.ModuleRouter{
	URLPrefix: "/api/account",
	SubRouters: []router.SubRouter{
		{
			Name:        "register",
			Methods:     []string{http.MethodPost},
			Pattern:     "/register",
			HandlerFunc: register,
		},
	},
}

type registerResponse struct {
	IsSuccess bool `json:"is_success"`
}

func register(w http.ResponseWriter, r *http.Request) {

	var requestBody struct{
		AccountName string `json:"account_name"`
		Password string `json:"password"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session:=mongo.CopySession()
	defer session.Close()



	var response registerResponse
	response.IsSuccess = true
	router.JSONResp(w, http.StatusOK, response)
}
