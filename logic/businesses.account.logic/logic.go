package businesses_account_logic

import (
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/config"
	"BuffetSalesManage/BuffetSalesManage/model/businesses.account.model"
	"gopkg.in/mgo.v2/bson"
	"BuffetSalesManage/BuffetSalesManage/base/error.code"
)

func IsExists(session *mgo.Session, phone string) bool {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	count, _ := coll.Find(bson.M{businesses_account_model.Phone.String(): phone}).Count()
	if count == 0 {
		return false
	}
	return true
}

func RegisterBusinesses(session *mgo.Session, phone, password string) error {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	selector := bson.M{businesses_account_model.Phone.String(): phone}
	update := bson.M{"$set": bson.M{businesses_account_model.Password.String(): password}}
	_, err := coll.Upsert(selector, update)
	return err
}

func IsLogin(session *mgo.Session, phone, password string) string {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)

	selector := bson.M{businesses_account_model.Phone.String(): phone}
	count, _ := coll.Find(selector).Count()
	if count == 0 {
		return ec.ACCOUNT_IS_NOT_EXISTS
	}
	var businessesInfo businesses_account_model.BusinessesAccount
	coll.Find(selector).One(&businessesInfo)
	if businessesInfo.Password != password {
		return ec.INVALID_PASSWORD
	}
	return ec.LOGIN_SUCCESS
}

func ComplementInfo(session *mgo.Session, businessesInfo businesses_account_model.BusinessesAccount) error {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	selector := bson.M{businesses_account_model.Phone.String(): businessesInfo.Phone}
	upset := bson.M{}
	if len(businessesInfo.NameOfShop) != 0 {
		upset[businesses_account_model.NameOfShop.String()] = businessesInfo.NameOfShop
	}
	if len(businessesInfo.Location) != 0 {
		upset[businesses_account_model.Location.String()] = businessesInfo.Location
	}
	if len(businessesInfo.Street) != 0 {
		upset[businesses_account_model.Street.String()] = businessesInfo.Street
	}
	if len(businessesInfo.PortraitOfShop) != 0 {
		upset[businesses_account_model.PortraitOfShop.String()] = businessesInfo.PortraitOfShop
	}
	update := bson.M{
		"$set": upset,
	}

	err := coll.Update(selector, update)
	return err
}
