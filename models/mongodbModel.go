package models

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var session *mgo.Session
var database *mgo.Database

const (
	MONGOSERVER = "localhost:27017"//这里使用了默认端口
	DATABASE = "vote_admin"//填入数据库名字

	MONGO_COLLECTION_QUESTION = "vote_question"
	MONGO_COLLECTION_VOTE_LOG = "vote_log"
)

func init(){
	//MongoDB
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MONGOSERVER},//这里填数据库地址
		Timeout:  60 * time.Second,
		Database: DATABASE,
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
	database = session.DB(DATABASE)//使用数据库
}
func GetMgo() *mgo.Session {
	return session
}
func GetDatabase() *mgo.Database {
	return database
}