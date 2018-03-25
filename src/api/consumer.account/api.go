package consumer_account

import (
	"BuffetSalesManage/BuffetSalesManage/router"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"BuffetSalesManage/BuffetSalesManage/base/error.code"
	"BuffetSalesManage/BuffetSalesManage/model/mongo"
	"BuffetSalesManage/BuffetSalesManage/model/consumer.account.model"
	"BuffetSalesManage/BuffetSalesManage/logic/consumer.account.logic"
)

var ExRouter = router.ModuleRouter{
	URLPrefix: "/api/consumer/account",
	SubRouters: []router.SubRouter{
		{
			Name:        "register",
			Methods:     []string{http.MethodPost},
			Pattern:     "/register",
			HandlerFunc: register,
		},
		{
			Name:        "EditInfo",
			Methods:     []string{http.MethodPost},
			Pattern:     "/info/complement",
			HandlerFunc: EditInfo,
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
		Phone    string `json:"phone"`
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

	session := mongo.CopySession()
	defer session.Close()

	isExists := consumer_account_logic.IsExists(session, requestBody.Phone)
	if isExists {
		router.JSONResp(w, http.StatusBadRequest, ec.AccountIsExists)
		return
	}

	err = consumer_account_logic.RegisterBusinesses(session, requestBody.Phone, requestBody.Password)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}

	router.JSONResp(w, http.StatusOK, nil)
}

func EditInfo(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Phone    string `json:"phone"`
		Location string `json:"location"`
		Portrait string `json:"portrait"`
		Nickname string `json:"nickname"`
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

	var consumerAccount consumer_account_model.ConsumerAccount
	consumerAccount.Phone = requestBody.Phone
	consumerAccount.Location = requestBody.Location
	consumerAccount.Portrait = requestBody.Portrait
	consumerAccount.Nickname = requestBody.Nickname

	err = consumer_account_logic.ComplementInfo(session, consumerAccount)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}
	router.JSONResp(w, http.StatusOK, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Phone    string `json:"phone"`
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

	session := mongo.CopySession()
	defer session.Close()

	isLogin := consumer_account_logic.IsLogin(session, requestBody.Phone, requestBody.Password)
	if !isLogin {
		router.JSONResp(w, http.StatusBadRequest, nil)
		return
	}
	router.JSONResp(w, http.StatusOK, nil)
}
