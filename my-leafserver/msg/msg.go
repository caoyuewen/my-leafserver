package msg

import "github.com/name5566/leaf/network/json"

var Processor = json.NewProcessor()

func init() {

	// REQUEST c - s
	Processor.Register(&CAuthorize{}) // 客户端认证服务端
	Processor.Register(&LoginReq{})   // 用户登录

	// RESPONSE
	// 返回结果 s - c
	Processor.Register(&Resp{}) // 返回结果

}
