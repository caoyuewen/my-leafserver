
// 客户端认证服务器
const C_AUTHORIZE_MSG_ID = 1; // 客户端验证服务器
const S_AUTHORIZE_MSG_ID = 2; // 服务器验证客户端

// 服务器返回的状态码
const successCode = 200;
const errCode = 500;


function CAuthorize(resp) {
    console.log("resp_data", resp);
    let result = resp["Resp"];
    let code = result["code"];
    let msg = result["msg"];

    if (code != successCode) {
        alert("Auth fail:" + msg);
        return
    }

    let data = result["data"];
    let sNonce = data["s_nonce"];
    let cNonce = data["c_nonce"];

    let decryptSNonce = DecryptNonce(sNonce);
    let decryptCNonce = DecryptNonce(cNonce);
    console.log("decryptCNonce", decryptCNonce);
    console.log("decryptSNonce", decryptSNonce);

    if (decryptCNonce != auth) {
        return;
    }
    // auth client
    let ensnonce = EncryptNonce(decryptSNonce);
    let encnonce = EncryptNonce(decryptCNonce);
    let sAuthorize = `{
		"SAuthorizeReq": {
		        "c_nonce":"` + encnonce + `",
		        "s_nonce":"` + ensnonce + `"
		    }
	    }`;
    doSend(sAuthorize)
}

function SAuthorize(resp) {
    console.log("resp_data", resp);
    let result = resp["Resp"];
    let code = result["code"];
    let msg = result["msg"];

    if (code != successCode) {
        alert("Auth fail:" + msg);
        return
    }
    alert("服务器验证成功我可以登陆了"+msg)

}