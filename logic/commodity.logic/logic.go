package commodity_logic

import (
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/model/commodity.model"
	"gopkg.in/mgo.v2/bson"
	"BuffetSalesManage/BuffetSalesManage/config"
)

func AddCommodity(session *mgo.Session, commodityInfo commodity_model.Commodity) error {
	insertInfo := bson.M{}
	if len(commodityInfo.CommodityType) != 0 {
		insertInfo[commodity_model.CommodityType.String()] = commodityInfo.CommodityType
	}
	insertInfo[commodity_model.BusinessId.String()] = commodityInfo.BusinessId
	if len(commodityInfo.CommodityName) != 0 {
		insertInfo[commodity_model.CommodityName.String()] = commodityInfo.CommodityName
	}
	if len(commodityInfo.CommodityPhoto) != 0 {
		insertInfo[commodity_model.CommodityPhoto.String()] = commodityInfo.CommodityPhoto
	}
	insertInfo[commodity_model.CommodityPrice.String()] = commodityInfo.CommodityPrice
	insertInfo[commodity_model.CommodityNum.String()] = commodityInfo.CommodityNum
	coll := session.DB(config.MongoDBName).C(commodity_model.COLL_COMMODITY)
	err := coll.Insert(insertInfo)
	return err
}

func Replenishment(session *mgo.Session, commodityId bson.ObjectId, addNum int) error {
	coll := session.DB(config.MongoDBName).C(commodity_model.COLL_COMMODITY)
	selector := bson.M{commodity_model.Id.String(): commodityId}
	var commodity commodity_model.Commodity
	coll.Find(selector).One(&commodity)
	commodityNum := commodity.CommodityNum + addNum
	err := coll.Update(selector, bson.M{"$set": bson.M{commodity_model.CommodityNum.String(): commodityNum}})
	return err
}

func GetCommodityBySelector(session *mgo.Session, selector bson.M) []commodity_model.Commodity {
	coll := session.DB(config.MongoDBName).C(commodity_model.COLL_COMMODITY)
	iter := coll.Find(selector).Iter()
	var (
		commodityList []commodity_model.Commodity
		item          commodity_model.Commodity
	)
	for iter.Next(&item) {
		commodityList = append(commodityList, item)
	}
	return commodityList
}

func IsEnoughCommodityNum(session *mgo.Session, commodityId bson.ObjectId, buyNum int) (bool, error) {
	coll := session.DB(config.MongoDBName).C(commodity_model.COLL_COMMODITY)
	selector := bson.M{commodity_model.Id.String(): commodityId}
	var commodity commodity_model.Commodity
	err := coll.Find(selector).One(&commodity)
	if err != nil {
		return false, err
	}
	commodityNum := commodity.CommodityNum - buyNum
	if commodityNum < 0 {
		return false, nil
	}
	return true, nil
}

func GetTotalMoney(session *mgo.Session, commodityId bson.ObjectId, buyNum int) int64 {
	selector := bson.M{commodity_model.Id.String(): commodityId}
	var commodity commodity_model.Commodity
	coll := session.DB(config.MongoDBName).C(commodity_model.COLL_COMMODITY)
	coll.Find(selector).One(&commodity)
	return commodity.CommodityPrice * int64(buyNum)
}

func GetNeedReplenishmentCommodity(session *mgo.Session, businessId bson.ObjectId) []string {
	coll := session.DB(config.MongoDBName).C(commodity_model.COLL_COMMODITY)
	selector := bson.M{commodity_model.BusinessId.String(): businessId}
	iter := coll.Find(selector).Iter()
	var (
		item                  commodity_model.Commodity
		needReplenishmentList []string
	)
	for iter.Next(&item) {
		if item.CommodityNum < 10 {
			needReplenishmentList = append(needReplenishmentList, item.CommodityName)
		}
	}
	return needReplenishmentList
}
