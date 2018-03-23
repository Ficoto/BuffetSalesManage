package consumer_account_model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage.git/config"
	"log"
)

const (
	COLL_CONSUMER_ACCOUNT = "consumer_account"
)

type ConsumerAccountKey int

const (
	Id          ConsumerAccountKey = iota
	AccountName
	Password
	Nickname
	Location
	Portrait
)

func (key ConsumerAccountKey) String() string {
	switch key {
	case Id:
		return "_id"
	case AccountName:
		return "account_name"
	case Password:
		return "password"
	case Nickname:
		return "nickname"
	case Location:
		return "location"
	case Portrait:
		return "portrait"
	default:
		return ""
	}
}

type ConsumerAccount struct {
	Id          bson.ObjectId `bson:"_id"`
	AccountName string        `bson:"account_name"`
	Password    string        `bson:"password"`
	Nickname    string        `bson:"nickname"`
	Location    string        `bson:"location"`
	Portrait    string        `bson:"portrait"`
}

func Index(session *mgo.Session) (*mgo.Collection, error) {
	coll := session.DB(config.MongoDBName).C(COLL_CONSUMER_ACCOUNT)

	indexes := []mgo.Index{
		{
			Key:        []string{AccountName.String()},
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
