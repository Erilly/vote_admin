package models

import "gopkg.in/mgo.v2"

//https://blog.csdn.net/wangshubo1989/article/details/75105397
func RegisterMongo(){
	session, err := mgo.Dial(url)

}