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
	"gopkg.in/mgo.v2/bson"
	"BuffetSalesManage/BuffetSalesManage/utils"
	"log"
	"fmt"
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
		{
			Name:        "GetBusinesses",
			Methods:     []string{http.MethodGet},
			Pattern:     "/list",
			HandlerFunc: GetBusinesses,
		},
	},
}

type registerResponse struct {
	BusinessId string `json:"business_id"`
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

	isExists := businesses_account_logic.IsExists(session, requestBody.Phone)
	if isExists {
		router.JSONResp(w, http.StatusBadRequest, ec.AccountIsExists)
		return
	}

	businessId, err := businesses_account_logic.RegisterBusinesses(session, requestBody.Phone, requestBody.Password)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}

	var response registerResponse
	response.BusinessId = businessId.Hex()
	fmt.Println(response.BusinessId)

	router.JSONResp(w, http.StatusOK, response)
}

func EditStoreInfo(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		BusinessId     string `json:"business_id"`
		Location       string `json:"location"`
		Street         string `json:"street"`
		NameOfShop     string `json:"name_of_shop"`
		PortraitOfShop string `json:"portrait_of_shop"`
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
	if !bson.IsObjectIdHex(requestBody.BusinessId) {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session := mongo.CopySession()
	defer session.Close()

	var businessesInfo businesses_account_model.BusinessesAccount
	businessesInfo.Id = bson.ObjectIdHex(requestBody.BusinessId)
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

	businessId, loginInfo := businesses_account_logic.IsLogin(session, requestBody.Phone, requestBody.Password)
	if loginInfo == ec.ACCOUNT_IS_NOT_EXISTS {
		router.JSONResp(w, http.StatusBadRequest, ec.AccountIsNotExists)
		return
	} else if loginInfo == ec.INVALID_PASSWORD {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidPassword)
		return
	}

	var response registerResponse
	response.BusinessId = businessId.Hex()

	router.JSONResp(w, http.StatusOK, response)
}

type BusinessInfo struct {
	BusinessId     string `json:"business_id"`
	Street         string `json:"street"`
	NameOfShop     string `json:"name_of_shop"`
	PortraitOfShop string `json:"portrait_of_shop"`
}

type BusinessList struct {
	BusinessInfoList []*BusinessInfo `json:"business_info_list"`
}

func GetBusinesses(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Location string `schema:"location"`
	}
	err := utils.NewSchemaDecoder().Decode(&requestBody, r.URL.Query())
	if err != nil {
		log.Println(r.URL.Path, err)
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session := mongo.CopySession()
	defer session.Close()

	selector := bson.M{businesses_account_model.Location.String(): requestBody.Location}
	businesses := businesses_account_logic.GetBusinessesBySelector(session, selector)

	var response BusinessList

	for _, item := range businesses {
		businessInfo := new(BusinessInfo)
		businessInfo.BusinessId = item.Id.Hex()
		businessInfo.NameOfShop = item.NameOfShop
		businessInfo.PortraitOfShop = item.PortraitOfShop
		businessInfo.Street = item.Street
		response.BusinessInfoList = append(response.BusinessInfoList, businessInfo)
	}
	router.JSONResp(w, http.StatusOK, response)
}
