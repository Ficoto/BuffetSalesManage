package order_logic

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"BuffetSalesManage/BuffetSalesManage/config"
	"BuffetSalesManage/BuffetSalesManage/model/order.model"
	"time"
	"BuffetSalesManage/BuffetSalesManage/model/consumer.account.model"
	"BuffetSalesManage/BuffetSalesManage/model/commodity.model"
)

func BuyCommodity(session *mgo.Session, BusinessId bson.ObjectId, CommodityId bson.ObjectId, CommodityNum int, totalOrderMoney int64, consumerId bson.ObjectId) {
	orderColl := session.DB(config.MongoDBName).C(order_model.COLL_ORDER)
	orderColl.Insert(bson.M{
		order_model.BusinessId.String():      BusinessId,
		order_model.CommodityId.String():     CommodityId,
		order_model.CommodityNum.String():    CommodityNum,
		order_model.OrderTotalPrice.String(): totalOrderMoney,
		order_model.CreateOrderTime.String(): time.Now(),
		order_model.ConsumerId.String():      consumerId,
	})

	consumerColl := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	consumerSelector := bson.M{consumer_account_model.Id.String(): consumerId}
	var consumer consumer_account_model.ConsumerAccount
	consumerColl.Find(consumerSelector).One(&consumer)
	consumerColl.Update(bson.M{consumer_account_model.Id.String(): consumer.Id}, bson.M{"$set": bson.M{consumer_account_model.Balance.String(): consumer.Balance - totalOrderMoney}})

	commodityColl := session.DB(config.MongoDBName).C(commodity_model.COLL_COMMODITY)
	commoditySelector := bson.M{commodity_model.Id.String(): CommodityId}
	var commodity commodity_model.Commodity
	commodityColl.Find(commoditySelector).One(&commodity)
	commodityColl.Update(bson.M{commodity_model.Id.String(): commodity.Id}, bson.M{"$set": bson.M{commodity_model.CommodityNum.String(): commodity.CommodityNum - CommodityNum}})
}

func GetYesterdayStatistics(session *mgo.Session, BusinessId bson.ObjectId) (int64, int) {
	t := time.Now()
	afterTime := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, time.Local)
	beforeTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	selector := bson.M{
		order_model.BusinessId.String():      BusinessId,
		order_model.CreateOrderTime.String(): bson.M{"$gte": afterTime, "$lt": beforeTime},
	}
	coll := session.DB(config.MongoDBName).C(order_model.COLL_ORDER)
	iter := coll.Find(selector).Iter()
	var (
		order             order_model.OrderInfo
		totalSales        int64
		salesCommodityNum int
	)
	for iter.Next(&order) {
		totalSales += order.OrderTotalPrice
		salesCommodityNum += order.CommodityNum
	}
	return totalSales, salesCommodityNum
}
