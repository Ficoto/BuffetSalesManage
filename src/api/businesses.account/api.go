package businesses_account

import (
	"BuffetSalesManage/BuffetSalesManage/router"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"BuffetSalesManage/BuffetSalesManage/base/error.code"
	"BuffetSalesManage/BuffetSalesManage/model/mongo"
	"BuffetSalesManage/BuffetSalesManage/logic/businesses.account.logic"
	"BuffetSalesManage/BuffetSalesManage/model/businesses.account.model"
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
		{
			Name:        "Login",
			Methods:     []string{http.MethodPost},
			Pattern:     "/login",
			HandlerFunc: Login,
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

func Login(w http.ResponseWriter, r *http.Request) {
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

	isLogin := businesses_account_logic.IsLogin(session, requestBody.AccountName, requestBody.Password)
	if !isLogin {
		router.JSONResp(w, http.StatusBadRequest, nil)
		return
	}
	router.JSONResp(w, http.StatusOK, nil)
}
