package mongo

import (
	"log"
	"sync"
	"gopkg.in/mgo.v2"
	"gitlab.xinghuolive.com/Backend-Go/orca/config"
)

var (
	session     *mgo.Session
	sessionOnce sync.Once
)

// Connect - Connect to mongo db
func Connect() {
	sessionOnce.Do(func() {
		var err error
		log.Printf("INFO(%s): Echo mongo conn uri.", config.MongoConnConf.URI)
		session, err = mgo.Dial(config.MongoConnConf.URI)
		if err != nil {
			log.Printf("FETAL(%s): Mongo connect error!", err.Error())
		} else {
			log.Println("SUCCESS(nil): Mongo connect success!")
		}
		// SetMode changes the consistency mode for the session.
		session.SetMode(mgo.Monotonic, true)
	})
}

// CloseSession - 注意这个函数只关闭全局的Session，不是复制的Session
func CloseSession() {
	session.Close()
}

//CopySession - 每次需要使用数据库时, 调用这个函数, 这个函数复制了一个连接的信息,并且在并发需要的时候会创建新的连接
//  一定要在复制出来的对象中调用defer session.Close() 否则连接被泄露了
func CopySession() *mgo.Session {
	return session.Copy()
}
