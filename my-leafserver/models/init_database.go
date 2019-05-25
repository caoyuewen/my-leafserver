package models

import (
	"github.com/wonderivan/logger"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	mgSession   *mgo.Session
	mgSessionDB *mgo.Database
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
	mgSession, err = mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		logger.Error("Mongodb 连接失败:", err)
		panic(err)
	}
	mgSession.SetMode(mgo.Monotonic, true)

	mgSessionDB = mgSession.DB(DBName)
	logger.Info("连接mongo数据库成功 address:", Server.MongoDBAddr)

}
