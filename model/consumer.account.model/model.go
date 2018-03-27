package consumer_account_model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/config"
	"log"
)

const (
	COLL_CONSUMER_ACCOUNT = "consumer_account"
)

type ConsumerAccountKey int

const (
	Id       ConsumerAccountKey = iota
	Phone
	Password
	Nickname
	Location
	Portrait
	Balance
)

func (key ConsumerAccountKey) String() string {
	switch key {
	case Id:
		return "_id"
	case Phone:
		return "phone"
	case Password:
		return "password"
	case Nickname:
		return "nickname"
	case Location:
		return "location"
	case Portrait:
		return "portrait"
	case Balance:
		return "balance"
	default:
		return ""
	}
}

type ConsumerAccount struct {
	Id       bson.ObjectId `bson:"_id"`
	Phone    string        `bson:"phone"`
	Password string        `bson:"password"`
	Nickname string        `bson:"nickname"`
	Location string        `bson:"location"`
	Portrait string        `bson:"portrait"`
	Balance  int64         `bson:"balance"`
}

func Index(session *mgo.Session) (*mgo.Collection, error) {
	coll := session.DB(config.MongoDBName).C(COLL_CONSUMER_ACCOUNT)

	indexes := []mgo.Index{
		{
			Key:        []string{Phone.String()},
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
