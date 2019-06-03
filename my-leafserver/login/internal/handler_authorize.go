package internal

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/name5566/leaf/gate"
	"github.com/wonderivan/logger"
	"my-leafserver/common"
	"my-leafserver/msg"
)

// 客户端认证服务端
func handlerCAuthorize(args []interface{}) {
	req := args[0].(*msg.CAuthorizeReq)
	currAgent := args[1].(gate.Agent)
	encNonce := req.CNonce
	resp := msg.NewSuccessResp(msg.CAuthorizeMsgId, "ok", nil)

	defer common.AccessRecords("handlerCAuthorize", req, &resp)

	// 使用base64解码 encNonce
	decodeString, err := base64.StdEncoding.DecodeString(encNonce)
	if err != nil {
		resp = msg.NewFailedResp(msg.CAuthorizeMsgId, "Decode fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}
	//encNonceBytes, err := common.RsaDecryptPri([]byte(encNonce))
	// 获取cNonce 明文
	encNonceBytes, err := common.RsaDecryptPri(decodeString)
	if err != nil {
		resp = msg.NewFailedResp(msg.CAuthorizeMsgId, "Decrypt fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	var data msg.CAuthorizeResp
	cNonceBytesEnc, err := common.RsaEncryptPub(encNonceBytes)
	if err != nil {
		resp = msg.NewFailedResp(msg.CAuthorizeMsgId, "Encrypt c fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	encryptCNonce := base64.StdEncoding.EncodeToString(cNonceBytesEnc)
	data.CNonce = encryptCNonce
	sNonce := uuid.New().String()
	data.SNonce, err = common.GetSNonce(sNonce)
	if err != nil {
		resp = msg.NewFailedResp(msg.CAuthorizeMsgId, "Encrypt s fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	logger.Info("cNonce:", string(encNonceBytes))
	logger.Info("sNonce:", sNonce)

	// 保存客户端
	saveAgent(string(encNonceBytes)+sNonce, &currAgent)

	resp.Data = data
	currAgent.WriteMsg(&resp)

}

// 服务端认证客户端
func handlerSAuthorize(args []interface{}) {
	req := args[0].(*msg.SAuthorizeReq)
	enCNonce := req.CNonce
	enSNonce := req.SNonce
	currAgent := args[1].(gate.Agent)


	resp := msg.NewSuccessResp(msg.SAuthorizeMsgId, "ok", nil)
	defer common.AccessRecords("handlerSAuthorize", req, &resp)

	enCNonceBytes, err := base64.StdEncoding.DecodeString(enCNonce)
	if err != nil {
		resp = msg.NewFailedResp(msg.SAuthorizeMsgId, "Decode fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	cNonceBytes, err := common.RsaDecryptPri(enCNonceBytes)
	if err != nil {
		resp = msg.NewFailedResp(msg.SAuthorizeMsgId, "Decrypt fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	enSNonceBytes, err := base64.StdEncoding.DecodeString(enSNonce)
	if err != nil {
		resp = msg.NewFailedResp(msg.SAuthorizeMsgId, "Decode fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	sNonceBytes, err := common.RsaDecryptPri(enSNonceBytes)
	if err != nil {
		resp = msg.NewFailedResp(msg.SAuthorizeMsgId, "Decrypt fail!", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	cNonce := string(cNonceBytes)
	sNonce := string(sNonceBytes)

	agent := getAgent(cNonce + sNonce)
	if agent == nil {
		resp = msg.NewFailedResp(msg.SAuthorizeMsgId, "Illegal auth or expired !", nil)
		currAgent.WriteMsg(&resp)
		return
	}

	(*agent).WriteMsg(&resp)
	currAgent.UserData()
	//currAgent.WriteMsg(&resp)
}
