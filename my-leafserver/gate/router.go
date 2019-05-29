package gate

import (
	"my-leafserver/login"
	"my-leafserver/msg"
)

func init() {
	// 客户端认证
	msg.Processor.SetRouter(&msg.CAuthorize{},login.ChanRPC)


	// 用户登录
	msg.Processor.SetRouter(&msg.LoginReq{}, login.ChanRPC)
}
