package msg

// RESPONSE OF ALL INCLUDE DATA

type Resp struct {
	MsgId int32       `json:"msg_id"`
	Code  int32       `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// RESPONSE OF ALL
//type Resp struct {
//	MsgId int32  `json:"msg_id"`
//	Code  int32  `json:"code"`
//	Msg   string `json:"msg"`
//}

func NewSuccessResp(msgId int32, msg string, data interface{}) Resp {
	return Resp{msgId, 200, msg, data}
}

func NewFailedResp(msgId int32, msg string, data interface{}) Resp {
	return Resp{msgId, 500, msg, data}
}
