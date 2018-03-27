package order_model

import (
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/config"
	"log"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	COLL_ORDER = "order"
)

type OrderKey int

const (
	Id              OrderKey = iota
	BusinessId
	CommodityId
	CommodityNum
	OrderTotalPrice
	CreateOrderTime
	ConsumerId
)

func (key OrderKey) String() string {
	switch key {
	case Id:
		return "_id"
	case BusinessId:
		return "business_id"
	case CommodityId:
		return "commodity_id"
	case CommodityNum:
		return "commodity_num"
	case OrderTotalPrice:
		return "order_total_price"
	case CreateOrderTime:
		return "create_order_time"
	case ConsumerId:
		return "consumer_id"
	default:
		return ""
	}
}

type OrderInfo struct {
	Id              bson.ObjectId `bson:"_id"`
	BusinessId      bson.ObjectId `bson:"business_id"`
	CommodityId     bson.ObjectId `bson:"commodity_id"`
	OrderTotalPrice int64         `bson:"order_total_price"`
	CreateOrderTime time.Time     `bson:"create_order_time"`
	CommodityNum    int           `bson:"commodity_num"`
	ConsumerId      bson.ObjectId `bson:"consumer_id"`
}

func Index(session *mgo.Session) (*mgo.Collection, error) {
	coll := session.DB(config.MongoDBName).C(COLL_ORDER)

	indexes := []mgo.Index{
		{
			Key:        []string{BusinessId.String()},
			Background: true,
		},
		{
			Key:        []string{CreateOrderTime.String()},
			Background: true,
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
