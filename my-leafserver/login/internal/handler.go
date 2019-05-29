package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"my-leafserver/common"
	"my-leafserver/msg"
	"reflect"
	"strconv"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.CAuthorize{}, handlerCAuthorize)
	handleMsg(&msg.LoginReq{}, handlerLogin)
}

// 用户认证
func handlerCAuthorize(args []interface{}) {
	authorize := args[0].(*msg.CAuthorize)
	agent := args[1].(gate.Agent)
	cNonce := authorize.CNonce
	resp := msg.NewSuccessResp(msg.CAuthorizeMsgId, "ok", nil)
	defer common.AccessRecords("handlerCAuthorize", authorize, &resp)

	var data msg.SAuthorize
	bytes, err := common.RsaDecrypt([]byte(cNonce), common.PriPerm)
	if err != nil {
		resp = msg.NewFailedResp(msg.CAuthorizeMsgId, "Decrypt Fail", err.Error())
		agent.WriteMsg(&resp)
		return
	}

	data.CNonce = string(bytes)
	data.SNonce = strconv.Itoa(common.RndNum(0, 20000))

	resp.Data = data
	agent.WriteMsg(&resp)

}

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
