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

func RegisterBusinesses(session *mgo.Session, phone, password string) (bson.ObjectId, error) {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	selector := bson.M{businesses_account_model.Phone.String(): phone}
	update := bson.M{"$set": bson.M{businesses_account_model.Password.String(): password}}
	_, err := coll.Upsert(selector, update)
	if err != nil {
		return bson.NewObjectId(), err
	}
	var business businesses_account_model.BusinessesAccount
	err = coll.Find(selector).One(&business)
	if err != nil {
		return bson.NewObjectId(), err
	}
	return business.Id, nil
}

func IsLogin(session *mgo.Session, phone, password string) (bson.ObjectId, string) {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)

	selector := bson.M{businesses_account_model.Phone.String(): phone}
	count, _ := coll.Find(selector).Count()
	if count == 0 {
		return bson.NewObjectId(), ec.ACCOUNT_IS_NOT_EXISTS
	}
	var businessesInfo businesses_account_model.BusinessesAccount
	coll.Find(selector).One(&businessesInfo)
	if businessesInfo.Password != password {
		return bson.NewObjectId(), ec.INVALID_PASSWORD
	}
	return businessesInfo.Id, ec.LOGIN_SUCCESS
}

func ComplementInfo(session *mgo.Session, businessesInfo businesses_account_model.BusinessesAccount) error {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	selector := bson.M{businesses_account_model.Id.String(): businessesInfo.Id}
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

func GetBusinessesBySelector(session *mgo.Session, selector bson.M) []businesses_account_model.BusinessesAccount {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	iter := coll.Find(selector).Iter()
	var (
		item       businesses_account_model.BusinessesAccount
		businesses []businesses_account_model.BusinessesAccount
	)
	for iter.Next(&item) {
		businesses = append(businesses, item)
	}
	return businesses
}

func GetBusinessInfo(session *mgo.Session, businessId bson.ObjectId) (businesses_account_model.BusinessesAccount, error) {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	selector := bson.M{businesses_account_model.Id.String(): businessId}
	var businessInfo businesses_account_model.BusinessesAccount
	err := coll.Find(selector).One(&businessInfo)
	return businessInfo, err
}
