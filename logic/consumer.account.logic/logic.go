package consumer_account_logic

import (
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/config"
	"gopkg.in/mgo.v2/bson"
	"BuffetSalesManage/BuffetSalesManage/model/consumer.account.model"
	"BuffetSalesManage/BuffetSalesManage/base/error.code"
)

func IsExists(session *mgo.Session, phone string) bool {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	count, _ := coll.Find(bson.M{consumer_account_model.Phone.String(): phone}).Count()
	if count != 0 {
		return false
	}
	return true
}

func RegisterBusinesses(session *mgo.Session, phone, password string) error {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	selector := bson.M{consumer_account_model.Phone.String(): phone}
	update := bson.M{"$set": bson.M{consumer_account_model.Password.String(): password}}
	_, err := coll.Upsert(selector, update)
	return err
}

func IsLogin(session *mgo.Session, phone, password string) string {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)

	selector := bson.M{consumer_account_model.Phone.String(): phone}
	var businessesInfo consumer_account_model.ConsumerAccount
	count, _ := coll.Find(selector).Count()
	if count == 0 {
		return ec.ACCOUNT_IS_NOT_EXISTS
	}
	coll.Find(selector).One(&businessesInfo)
	if businessesInfo.Password != password {
		return ec.INVALID_PASSWORD
	}
	return ec.LOGIN_SUCCESS
}

func ComplementInfo(session *mgo.Session, consumerAccount consumer_account_model.ConsumerAccount) error {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	selector := bson.M{consumer_account_model.Phone.String(): consumerAccount.Phone}
	update := bson.M{
		"$set": bson.M{
			consumer_account_model.Nickname.String(): consumerAccount.Nickname,
			consumer_account_model.Location.String(): consumerAccount.Location,
			consumer_account_model.Portrait.String(): consumerAccount.Portrait,
		},
	}

	err := coll.Update(selector, update)
	return err
}
