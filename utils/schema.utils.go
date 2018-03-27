package utils

import (
	"github.com/gorilla/schema"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

func convertBsonObjectId(value string) reflect.Value {
	if bson.IsObjectIdHex(value) {
		return reflect.ValueOf(bson.ObjectIdHex(value))
	} else {
		return reflect.Value{}
	}
}

func NewSchemaDecoder() *schema.Decoder {
	schemaDecoder := schema.NewDecoder()
	schemaDecoder.IgnoreUnknownKeys(true)
	schemaDecoder.RegisterConverter(bson.NewObjectId(), convertBsonObjectId)
	return schemaDecoder
}
