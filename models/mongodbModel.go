package models

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
	"github.com/astaxie/beego"
)

var session *mgo.Session
var database *mgo.Database

const (
	MONGO_COLLECTION_QUESTION = "vote_question"
	MONGO_COLLECTION_VOTE_LOG = "vote_log"
)

func init(){
	//MongoDB
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{beego.AppConfig.String("mongo_server")},//这里填数据库地址
		Timeout:  60 * time.Second,
		Database: beego.AppConfig.String("mongo_dbname"),
		PoolLimit: 4096,//
		//Username: AuthUserName,
		//Password: AuthPassword,
	}
	var err error
	session, err = mgo.DialWithInfo(mongoDBDialInfo)//连接数据库
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)
	database = session.DB(beego.AppConfig.String("mongo_dbname"))//使用数据库
	//mgo.SetDebug(true)  // 设置DEBUG模式
	//mgo.SetLogger(new(MongoLog)) // 设置日志.
}
func GetMgo() *mgo.Session {
	return session
}
func GetDatabase() *mgo.Database {
	return database
}

// 实现 mongo.Logger 的接口
//type MongoLog struct {
//}
//
//func (MongoLog)Output(calldepth int, s string) error {
//	log.SetFlags(log.Lshortfile)
//	return log.Output(calldepth,s)
//}
