package msg

// REQUEST
// 客户端认证服务端请求
type CAuthorizeReq struct {
	CNonce string `json:"c_nonce"`
}

// RESPONSE DATA  MSG_ID:1
// 客户端认证服务端请求返回结果
type CAuthorizeResp struct {
	CNonce string `json:"c_nonce"`
	SNonce string `json:"s_nonce"`
}

// REQUEST
// 客户端认证服务端请求
type SAuthorizeReq struct {
	CNonce string `json:"c_nonce"`
	SNonce string `json:"s_nonce"`
}

// RESPONSE DATA MSG_ID:2
// 客户端认证服务端请求
// RESP OF ALL
/*
	{
		msg_id:2,

	}
*/


// 登录请求
type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
