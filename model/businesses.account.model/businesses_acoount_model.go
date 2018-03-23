package businesses_account_model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/config"
	"log"
)

const (
	COLL_BUSINESSES_ACCOUNT = "businesses_account"
)

type BusinessesAccountKey int

const (
	Id          BusinessesAccountKey = iota
	AccountName
	Password
	NameOfShop
	Location
	Street
)

func (key BusinessesAccountKey) String() string {
	switch key {
	case Id:
		return "_id"
	case AccountName:
		return "account_name"
	case Password:
		return "password"
	case NameOfShop:
		return "name_of_shop"
	case Location:
		return "location"
	case Street:
		return "street"
	default:
		return ""
	}
}

type BusinessesAccount struct {
	Id          bson.ObjectId `bson:"_id"`
	AccountName string        `bson:"account_name"`
	Password    string        `bson:"password"`
	NameOfShop  string        `bson:"name_of_shop"`
	Location    string        `json:"location"`
	Street      string        `bson:"street"`
}

func Index(session *mgo.Session) (*mgo.Collection, error) {
	coll := session.DB(config.MongoDBName).C(COLL_BUSINESSES_ACCOUNT)

	indexes := []mgo.Index{
		{
			Key:        []string{AccountName.String()},
			Background: true,
		},
		{
			Key:        []string{Location.String()},
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
