package models

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"time"
)

var (
	// gate conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
)

var Server struct {
	WSAddr     string
	TCPAddr    string
	MaxConnNum int

	//LogLevel    string
	LogPath     string
	CertFile    string
	KeyFile     string
	ConsolePort int
	ProfilePath string

	MongoDBAddr string
	MongoDBUser string
	MongoDBPwd  string
	MongoDBAuth string
}

func init() {
	initLogger()
	initServerConf()

	// 初始化数据库
	InitDB()

	var user User
	user.Username = "test"
	user.Password = "123456"
	err := user.AddUser()
	if err != nil {
		logger.Error("插入数据失败...!", err.Error())
		return
	}

	logger.Info("服务启动成功...")
}

// 加载日志适配器
func initLogger() {
	err := logger.SetLogger("conf/conf_log.json")
	if err != nil {
		panic("[conf/init_database.go] 加载日志配置失败 " + err.Error())
	}
}

// 加载服务器配置
func initServerConf() {
	data, err := ioutil.ReadFile("conf/conf_server.json")
	if err != nil {
		logger.Error("加载服务器配置失败!", err.Error())
		panic(err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		logger.Error("加载服务器配置失败!", err.Error())
		panic(err)
	}
	logger.Info("加载服务器配置成功,正在启动服务...")
}
