package order

import (
	"BuffetSalesManage/BuffetSalesManage/router"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"BuffetSalesManage/BuffetSalesManage/base/error.code"
	"gopkg.in/mgo.v2/bson"
	"BuffetSalesManage/BuffetSalesManage/model/mongo"
	"BuffetSalesManage/BuffetSalesManage/logic/commodity.logic"
	"BuffetSalesManage/BuffetSalesManage/logic/consumer.account.logic"
	"BuffetSalesManage/BuffetSalesManage/logic/order.logic"
	"BuffetSalesManage/BuffetSalesManage/utils"
	"log"
)

var ExRouter = router.ModuleRouter{
	URLPrefix: "/api/order",
	SubRouters: []router.SubRouter{
		{
			Name:        "BuyCommodity",
			Methods:     []string{http.MethodPost},
			Pattern:     "/buy",
			HandlerFunc: BuyCommodity,
		},
		{
			Name:        "GetStatistics",
			Methods:     []string{http.MethodGet},
			Pattern:     "/statistics",
			HandlerFunc: GetStatistics,
		},
	},
}

func BuyCommodity(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		BusinessId  string `json:"business_id"`
		ConsumerId  string `json:"consumer_id"`
		CommodityId string `json:"commodity_id"`
		BuyNum      int    `json:"buy_num"`
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
	if !bson.IsObjectIdHex(requestBody.BusinessId) || !bson.IsObjectIdHex(requestBody.ConsumerId) || !bson.IsObjectIdHex(requestBody.CommodityId) {
		router.JSONResp(w, http.StatusBadRequest, ec.InvalidArgument)
		return
	}

	session := mongo.CopySession()
	defer session.Close()

	totalMoney := commodity_logic.GetTotalMoney(session, bson.ObjectIdHex(requestBody.CommodityId), requestBody.BuyNum)
	isEnoughCommodityNum, err := commodity_logic.IsEnoughCommodityNum(session, bson.ObjectIdHex(requestBody.CommodityId), requestBody.BuyNum)
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}
	if !isEnoughCommodityNum {
		router.JSONResp(w, http.StatusBadRequest, ec.IsNotEnoughCommodity)
		return
	}
	isEnoughBalance, err := consumer_account_logic.IsEnoughBalance(session, totalMoney, bson.ObjectIdHex(requestBody.ConsumerId))
	if err != nil {
		router.JSONResp(w, http.StatusBadRequest, ec.MongodbOp)
		return
	}
	if !isEnoughBalance {
		router.JSONResp(w, http.StatusBadRequest, ec.IsNotEnoughBalance)
		return
	}
	order_logic.BuyCommodity(session, bson.ObjectIdHex(requestBody.BusinessId), bson.ObjectIdHex(requestBody.CommodityId), requestBody.BuyNum, totalMoney, bson.ObjectIdHex(requestBody.ConsumerId))
	router.JSONResp(w, http.StatusOK, nil)
}

type StatisticsInfo struct {
	TotalSalesMoney            int64    `json:"total_sales_money"`
	SalesCommodityNum          int      `json:"sales_commodity_num"`
	NeedReplenishmentCommodity []string `json:"need_replenishment_commodity"`
}

func GetStatistics(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		BusinessId string `schema:"business_id"`
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

	var response StatisticsInfo
	response.TotalSalesMoney, response.SalesCommodityNum = order_logic.GetYesterdayStatistics(session, bson.ObjectIdHex(requestBody.BusinessId))
	response.NeedReplenishmentCommodity = commodity_logic.GetNeedReplenishmentCommodity(session, bson.ObjectIdHex(requestBody.BusinessId))
	router.JSONResp(w, http.StatusOK, response)
}
