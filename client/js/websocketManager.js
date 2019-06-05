var wsUri = "ws://127.0.0.1:3666";

function init() {
    // output = document.getElementById("output");
    myWebSocket();
}

var auth;

function myWebSocket() {
    websocket = new WebSocket(wsUri);
    websocket.onopen = function (evt) {
        onOpen(evt)
    };
    websocket.onclose = function (evt) {
        onClose(evt)
    };
    websocket.onmessage = function (evt) {
        onMessage(evt)
    };
    websocket.onerror = function (evt) {
        onError(evt)
    };
}

function onOpen(evt) {
    auth = RndNum(100000) + 100000;

    //alert(auth);
    alert("建立连接,auth=" + auth);
    let nonce = EncryptNonce(auth.toString());
    let cAuthorize = CAuthorize;
    cAuthorize.cnonce = nonce;

    doSend(0, cAuthorize);
}

function onClose(evt) {
    alert("关闭连接");

}

function onMessage(evt) {
    //writeToScreen('<span style="color: blue;">RESPONSE: ' + evt.data + '</span>');
    let data = evt.data;
    // websocket.close();
    let reader = new FileReader();
    reader.addEventListener("loadend", function () {
        let msgObj = JSON.parse(reader.result);
        let msgId = msgObj["Resp"]["msg_id"];
        //console.log(msgObj);
        switch (msgId) {
            case C_AUTHORIZE_MSG_ID:
                CAuthorize(msgObj);
                break;
            case S_AUTHORIZE_MSG_ID:
                SAuthorize(msgObj);
                break;


            default:
                console.log("未知消息")
        }

    });
    reader.readAsBinaryString(data)
}

function onError(evt) {
    alert("onError");
    //writeToScreen('<span style="color: red;">ERROR:</span> ' + evt.data);
}

function doSend(msgId, message) {


    websocket.send();
}





