package internal

import (
	"my-leafserver/msg"
	"reflect"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.CAuthorizeReq{}, handlerCAuthorize)
	handleMsg(&msg.SAuthorizeReq{}, handlerSAuthorize)
	handleMsg(&msg.LoginReq{}, handlerLogin)
}

