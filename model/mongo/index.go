package mongo

import (
	"fmt"
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/model/businesses.account.model"
)

type indexFunc func(session *mgo.Session) (*mgo.Collection, error)

func createIndex(session *mgo.Session, fn indexFunc) {
	coll, err := fn(session)
	fmt.Printf("-----Index of %s: \r\n", coll.Name)
	if err != nil {
		log.Printf("ERROR(%s): Initialize index error!", err.Error())
	}
	indexList, _ := coll.Indexes()
	for _, index := range indexList {
		fmt.Printf("key: %+v ", index.Key)
		fmt.Printf("name: %s ", index.Name)
		fmt.Printf("unique: %t\r\n", index.Unique)
	}
}

// InitializeIndex - 初始化Index
func InitializeIndex() {
	session := CopySession()
	defer session.Close()

	startTime := time.Now()

	createIndex(session, businesses_account_model.Index)

	elapsed := time.Since(startTime)
	fmt.Printf("-----Create index elapsed: %s\r\n", elapsed)
}
