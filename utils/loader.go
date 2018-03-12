package utils

import (
	"github.com/spf13/viper"
	"sync"
)

//
var (
	mongoConf mongoConn
)

var (
	viperTemp *viper.Viper
	viperOnce sync.Once
)

func loadViper() *viper.Viper {
	viperOnce.Do(func() {
		viperTemp = LoadConfUtil(DBServerFile)
	})
	return viperTemp
}

// Mongo struct initialization
func LoadMongoConf() *mongoConn {
	conn := Conn{}
	v := loadViper()
	conn.setHostList(v.GetStringSlice("mongodb.hostport"))
	conn.generateURI()
	ac := authConn{
		conn:     conn,
		username: loadViper().GetString("mongodb.username"),
		password: loadViper().GetString("mongodb.password"),
	}
	ac.generateURI()
	mongoConf.authConn = ac
	mongoConf.AuthSource = loadViper().GetString("mongodb.auth_source")
	mongoConf.ReplicateSet = loadViper().GetString("mongodb.ReplicateSet")
	mongoConf.generateURI()

	return &mongoConf
}