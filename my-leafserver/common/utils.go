package common

import (
	"github.com/wonderivan/logger"
	"math/rand"
	"time"
)

var count int64
// 获取一个随机数
func RndNum(startNum, endNum int) int {
	count++
	if count >= 1<<8 {
		count = 0
	}
	rand.Seed(time.Now().UnixNano() + count)
	rnd := rand.Intn(endNum - startNum)
	return rnd + startNum
}



// 访问记录到控制台
func AccessRecords(acc string, req, resp interface{}) {
	logger.Info("")
	logger.Info(time.Now().Format(TimeFormat), ":", acc)
	logger.Info("req:%+v", req)
	logger.Info("rsp:%+v", resp)
	logger.Info("")
}
