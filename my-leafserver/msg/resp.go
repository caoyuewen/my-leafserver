package msg

/*--------------  RESP --------------------*/
type Resp struct {
	MsgId int32       `json:"msg_id"`
	Code  int32       `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func NewSuccessResp(msgId int32, msg string, data interface{}) Resp {
	return Resp{msgId, 200, msg, data}
}

func NewFailedResp(msgId int32, msg string, data interface{}) Resp {
	return Resp{msgId, 500, msg, data}
}
/*--------------  RESP --------------------*/

/*==============  DATA =====================*/
type SAuthorize struct {
	CNonce string `json:"c_nonce"`
	SNonce string `json:"s_nonce"`
}
/*==============  DATA =====================*/