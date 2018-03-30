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
	"gopkg.in/mgo.v2/bson"
	"fmt"
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
		{
			Name:        "RechargeBalance",
			Methods:     []string{http.MethodPost},
			Pattern:     "/balance/recharge",
			HandlerFunc: RechargeBalance,
		},
	},
}

type registerResponse struct {
	ConsumerId string `json:"consumer_id"`
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
	fmt.Println(isExists)
	if isExists {
		router.JSONResp(w, http.StatusBadRequest, ec.AccountIsExists)
		return
	}

	consumerId, err := consumer_account_logic.RegisterBusinesses(session, requestBody.Phone, requestBody.Password)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}

	var response registerResponse
	response.ConsumerId = consumerId.Hex()

	router.JSONResp(w, http.StatusOK, response)
}

func EditInfo(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		ConsumerId string `json:"consumer_id"`
		Location   string `json:"location"`
		Portrait   string `json:"portrait"`
		Nickname   string `json:"nickname"`
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
	if !bson.IsObjectIdHex(requestBody.ConsumerId) {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session := mongo.CopySession()
	defer session.Close()

	var consumerAccount consumer_account_model.ConsumerAccount
	consumerAccount.Id = bson.ObjectIdHex(requestBody.ConsumerId)
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

	consumerId, loginInfo := consumer_account_logic.IsLogin(session, requestBody.Phone, requestBody.Password)
	if loginInfo == ec.ACCOUNT_IS_NOT_EXISTS {
		router.JSONResp(w, http.StatusBadRequest, ec.AccountIsNotExists)
		return
	} else if loginInfo == ec.INVALID_PASSWORD {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidPassword)
		return
	}
	var response registerResponse
	response.ConsumerId = consumerId.Hex()

	router.JSONResp(w, http.StatusOK, response)
}

func RechargeBalance(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		ConsumerId    string `json:"consumer_id"`
		RechargeMoney int64  `json:"recharge_money"`
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
	if !bson.IsObjectIdHex(requestBody.ConsumerId) {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session := mongo.CopySession()
	defer session.Close()

	err = consumer_account_logic.RechargeBalance(session, requestBody.RechargeMoney, bson.ObjectIdHex(requestBody.ConsumerId))
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}
	router.JSONResp(w, http.StatusOK, nil)
}
