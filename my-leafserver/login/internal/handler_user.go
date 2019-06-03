package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"my-leafserver/msg"
)

// 用户登录
func handlerLogin(args []interface{}) {
	loginReq := args[0].(*msg.LoginReq)
	agent := args[1].(gate.Agent)

	account := loginReq.Account
	password := loginReq.Password
	fmt.Println(account, password)

	var resp msg.Resp
	defer fmt.Println("登录返回结果：", &resp)

	if account == "test" && password == "123456" {
		resp = msg.NewSuccessResp(msg.LoginReqMsgId, "登录成功", nil)
		agent.WriteMsg(&resp)
	} else {
		resp = msg.NewFailedResp(msg.LoginReqMsgId, "登录失败", nil)
		agent.WriteMsg(&resp)
	}
}

