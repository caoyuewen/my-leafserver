package models

import (
	"github.com/wonderivan/logger"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	globalSession *mgo.Session
)

const (
	DBName = "game"
)

// 初始化数据库
func InitDB() {
	initMongoDb()
}

// 初始化mongoDB
func initMongoDb() {
	//此处连接正式线上数据库  下面是模拟的直接连
	var err error
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{Server.MongoDBAddr},
		Timeout:  60 * time.Second,
		Database: Server.MongoDBAuth,
		Username: Server.MongoDBUser,
		Password: Server.MongoDBPwd,
	}
	globalSession, err = mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		logger.Error("Mongodb 连接失败:", err)
		panic(err)
	}
	globalSession.SetMode(mgo.Monotonic, true)
	logger.Info("连接mongo数据库成功 address:", Server.MongoDBAddr)
}

func connection(dbName, cName string) (*mgo.Session, *mgo.Collection) {
	s := globalSession.Copy()
	c := s.DB(dbName).C(cName)
	return s, c
}
