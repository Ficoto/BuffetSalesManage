package commodity_model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/config"
	"log"
)

const (
	COLL_COMMODITY = "commodity"
)

type CommodityKey int

const (
	Id             CommodityKey = iota
	BusinessId
	CommodityType
	CommodityName
	CommodityPhoto
	CommodityNum
	CommodityPrice
)

func (key CommodityKey) String() string {
	switch key {
	case Id:
		return "_id"
	case BusinessId:
		return "business_id"
	case CommodityType:
		return "commodity_type"
	case CommodityName:
		return "commodity_name"
	case CommodityPhoto:
		return "commodity_photo"
	case CommodityNum:
		return "commodity_num"
	case CommodityPrice:
		return "commodity_price"
	default:
		return ""
	}
}

type Commodity struct {
	Id             bson.ObjectId `bson:"_id"`
	BusinessId     bson.ObjectId `bson:"business_id"`
	CommodityType  string        `bson:"commodity_type"`
	CommodityName  string        `bson:"commodity_name"`
	CommodityPhoto string        `bson:"commodity_photo"`
	CommodityNum   int           `bson:"commodity_num"`
	CommodityPrice int64         `bson:"commodity_price"`
}

func Index(session *mgo.Session) (*mgo.Collection,error) {
	coll:=session.DB(config.MongoDBName).C(COLL_COMMODITY)

	indexes:=[]mgo.Index{
		{
			Key:[]string{BusinessId.String()},
			Background:true,
		},
		{
			Key:[]string{CommodityType.String()},
			Background:true,
		},
	}

	var lastErr error
	for _, index := range indexes {
		err := coll.EnsureIndex(index)
		if err != nil {
			lastErr = err
			log.Panicln("ensure index error", err)
		}
	}
	return coll, lastErr
}
