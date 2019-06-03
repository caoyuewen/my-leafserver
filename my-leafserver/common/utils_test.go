package common

import (
	"fmt"
	"testing"
)

func TestRndNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		num := RndNum(5, 60)
		fmt.Println(num)
	}

}

func TestRsaEncrypt(t *testing.T) {

	var oData = "8274"

	data, err := RsaEncryptPub([]byte(oData))
	if err != nil {
		fmt.Println("加密失败err：",err.Error())
		return
	}
	fmt.Println(string(data))

	bytes, err := RsaDecryptPri(data)
	if err != nil {
		fmt.Println("解密失败err：",err.Error())
		return
	}

	fmt.Println("解密后的数据：",string(bytes))

}