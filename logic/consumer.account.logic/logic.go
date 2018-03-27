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

func RegisterBusinesses(session *mgo.Session, phone, password string) (bson.ObjectId, error) {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	selector := bson.M{consumer_account_model.Phone.String(): phone}
	update := bson.M{"$set": bson.M{consumer_account_model.Password.String(): password}}
	_, err := coll.Upsert(selector, update)
	if err != nil {
		return bson.NewObjectId(), err
	}
	var consumer consumer_account_model.ConsumerAccount
	err = coll.Find(selector).One(&consumer)
	if err != nil {
		return bson.NewObjectId(), err
	}
	return consumer.Id, nil
}

func IsLogin(session *mgo.Session, phone, password string) (bson.ObjectId, string) {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)

	selector := bson.M{consumer_account_model.Phone.String(): phone}
	var businessesInfo consumer_account_model.ConsumerAccount
	count, _ := coll.Find(selector).Count()
	if count == 0 {
		return bson.NewObjectId(), ec.ACCOUNT_IS_NOT_EXISTS
	}
	coll.Find(selector).One(&businessesInfo)
	if businessesInfo.Password != password {
		return bson.NewObjectId(), ec.INVALID_PASSWORD
	}
	return businessesInfo.Id, ec.LOGIN_SUCCESS
}

func ComplementInfo(session *mgo.Session, consumerAccount consumer_account_model.ConsumerAccount) error {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	selector := bson.M{consumer_account_model.Id.String(): consumerAccount.Id}
	upset := bson.M{}
	if len(consumerAccount.Nickname) != 0 {
		upset[consumer_account_model.Nickname.String()] = consumerAccount.Nickname
	}
	if len(consumerAccount.Location) != 0 {
		upset[consumer_account_model.Location.String()] = consumerAccount.Location
	}
	if len(consumerAccount.Portrait) != 0 {
		upset[consumer_account_model.Portrait.String()] = consumerAccount.Portrait
	}
	update := bson.M{
		"$set": upset,
	}

	err := coll.Update(selector, update)
	return err
}

func RechargeBalance(session *mgo.Session, balance int64, consumerId bson.ObjectId) error {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	selector := bson.M{consumer_account_model.Id.String(): consumerId}
	var consumer consumer_account_model.ConsumerAccount
	err := coll.Find(selector).One(&consumer)
	if err != nil {
		return err
	}
	totalBalance := consumer.Balance + balance
	err = coll.Update(selector, bson.M{"$set": bson.M{consumer_account_model.Balance.String(): totalBalance}})
	return err
}

func IsEnoughBalance(session *mgo.Session, totalOrderMoney int64, consumerId bson.ObjectId) (bool, error) {
	coll := session.DB(config.MongoDBName).C(consumer_account_model.COLL_CONSUMER_ACCOUNT)
	selector := bson.M{consumer_account_model.Id.String(): consumerId}
	var consumer consumer_account_model.ConsumerAccount
	err := coll.Find(selector).One(&consumer)
	if err != nil {
		return false, err
	}
	if totalOrderMoney > consumer.Balance {
		return false, nil
	}
	return true, nil
}
