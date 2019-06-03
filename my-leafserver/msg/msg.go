package msg

import "github.com/name5566/leaf/network/json"

var Processor = json.NewProcessor()

func init() {

	// REQUEST  S - C
	Processor.Register(&CAuthorizeReq{}) // 客户端认证服务端
	Processor.Register(&SAuthorizeReq{}) // 服务端认证客户端

	Processor.Register(&LoginReq{}) // 用户登录

	// RESPONSE OF ALL S - C
	Processor.Register(&Resp{}) // 返回结果

}
