package mongo_test

import (
	"gitlab.xinghuolive.com/Backend-Go/kangaroo/model/mongo"
	"testing"
)

func TestInitializeIndex(t *testing.T) {
	mongo.Connect()
	defer mongo.CloseSession()
	mongo.InitializeIndex()
}
