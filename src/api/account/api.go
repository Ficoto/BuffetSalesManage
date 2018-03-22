package account

import (
	"BuffetSalesManage/BuffetSalesManage.git/router"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"BuffetSalesManage/BuffetSalesManage.git/base/error.code"
	"BuffetSalesManage/BuffetSalesManage.git/model/mongo"
	"BuffetSalesManage/BuffetSalesManage.git/logic/businesses.account.logic"
	"BuffetSalesManage/BuffetSalesManage.git/model/businesses.account.model"
)

var ExRouter = router.ModuleRouter{
	URLPrefix: "/api/business/account",
	SubRouters: []router.SubRouter{
		{
			Name:        "register",
			Methods:     []string{http.MethodPost},
			Pattern:     "/register",
			HandlerFunc: register,
		},
		{
			Name:        "EditStoreInfo",
			Methods:     []string{http.MethodPost},
			Pattern:     "/info/complement",
			HandlerFunc: EditStoreInfo,
		},
	},
}

func register(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		AccountName string `json:"account_name"`
		Password    string `json:"password"`
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

	session := mongo.CopySession()
	defer session.Close()

	isExists := businesses_account_logic.IsExists(session, requestBody.AccountName)
	if isExists {
		router.JSONResp(w, http.StatusBadRequest, ec.AccountIsExists)
		return
	}

	err = businesses_account_logic.RegisterBusinesses(session, requestBody.AccountName, requestBody.Password)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}

	router.JSONResp(w, http.StatusOK, nil)
}

func EditStoreInfo(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		AccountName string `json:"account_name"`
		Location    string `json:"location"`
		Street      string `json:"street"`
		NameOfShop  string `json:"name_of_shop"`
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

	session := mongo.CopySession()
	defer session.Close()

	var businessesInfo businesses_account_model.BusinessesAccount
	businessesInfo.AccountName = requestBody.AccountName
	businessesInfo.Location = requestBody.Location
	businessesInfo.Street = requestBody.Street
	businessesInfo.NameOfShop = requestBody.NameOfShop

	err = businesses_account_logic.ComplementInfo(session, businessesInfo)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}
	router.JSONResp(w, http.StatusOK, nil)
}
