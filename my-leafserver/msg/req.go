package msg

// 认证请求
type CAuthorize struct {
	CNonce string `json:"c_nonce"`
}

// 登录请求
type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
