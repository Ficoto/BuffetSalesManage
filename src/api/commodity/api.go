package commodity

import (
	"BuffetSalesManage/BuffetSalesManage/router"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"BuffetSalesManage/BuffetSalesManage/base/error.code"
	"BuffetSalesManage/BuffetSalesManage/model/commodity.model"
	"gopkg.in/mgo.v2/bson"
	"BuffetSalesManage/BuffetSalesManage/logic/commodity.logic"
	"BuffetSalesManage/BuffetSalesManage/model/mongo"
	"BuffetSalesManage/BuffetSalesManage/utils"
	"log"
)

var ExRouter = router.ModuleRouter{
	URLPrefix: "/api/commodity",
	SubRouters: []router.SubRouter{
		{
			Name:        "AddCommodity",
			Methods:     []string{http.MethodPost},
			Pattern:     "/add",
			HandlerFunc: AddCommodity,
		},
		{
			Name:        "Replenishment",
			Methods:     []string{http.MethodPost},
			Pattern:     "/replenishment",
			HandlerFunc: Replenishment,
		},
		{
			Name:        "GetCommodityList",
			Methods:     []string{http.MethodGet},
			Pattern:     "/list",
			HandlerFunc: GetCommodityList,
		},
	},
}

func AddCommodity(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		BusinessId     string `json:"business_id"`
		CommodityType  string `json:"commodity_type"`
		CommodityName  string `json:"commodity_name"`
		CommodityPhoto string `json:"commodity_photo"`
		CommodityPrice int64  `json:"commodity_price"`
		CommodityNum   int    `json:"commodity_num"`
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

	commodity := commodity_model.Commodity{
		BusinessId:     bson.ObjectIdHex(requestBody.BusinessId),
		CommodityType:  requestBody.CommodityType,
		CommodityName:  requestBody.CommodityName,
		CommodityNum:   requestBody.CommodityNum,
		CommodityPrice: requestBody.CommodityPrice,
		CommodityPhoto: requestBody.CommodityPhoto,
	}
	err = commodity_logic.AddCommodity(session, commodity)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}
	router.JSONResp(w, http.StatusOK, nil)
}

func Replenishment(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		CommodityId string `json:"commodity_id"`
		AddNum      int    `json:"add_num"`
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
	if !bson.IsObjectIdHex(requestBody.CommodityId) {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session := mongo.CopySession()
	defer session.Close()

	err = commodity_logic.Replenishment(session, bson.ObjectIdHex(requestBody.CommodityId), requestBody.AddNum)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}
	router.JSONResp(w, http.StatusOK, nil)
}

type CommodityInfo struct {
	CommodityId    string `json:"commodity_id"`
	CommodityName  string `json:"commodity_name"`
	CommodityPhoto string `json:"commodity_photo"`
	CommodityNum   int    `json:"commodity_num"`
	CommodityPrice int64  `json:"commodity_price"`
}

type CommodityInfoList struct {
	CommodityInfoes []*CommodityInfo `json:"commodity_infoes"`
}

func GetCommodityList(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		BusinessId    string `schema:"business_id"`
		CommodityType string `schema:"commodity_type"`
	}
	err := utils.NewSchemaDecoder().Decode(&requestBody, r.URL.Query())
	if err != nil {
		log.Println(r.URL.Path, err)
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}
	if !bson.IsObjectIdHex(requestBody.BusinessId) {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session := mongo.CopySession()
	defer session.Close()

	selector := bson.M{commodity_model.BusinessId.String(): bson.ObjectIdHex(requestBody.BusinessId)}
	if len(requestBody.CommodityType) != 0 {
		selector[commodity_model.CommodityType.String()] = requestBody.CommodityType
	}
	commodityList := commodity_logic.GetCommodityBySelector(session, selector)

	var response CommodityInfoList

	for _, item := range commodityList {
		info := new(CommodityInfo)
		info.CommodityId = item.Id.Hex()
		info.CommodityPhoto = item.CommodityPhoto
		info.CommodityPrice = item.CommodityPrice
		info.CommodityNum = item.CommodityNum
		info.CommodityName = item.CommodityName
		response.CommodityInfoes = append(response.CommodityInfoes, info)
	}
	router.JSONResp(w, http.StatusOK, response)
}
