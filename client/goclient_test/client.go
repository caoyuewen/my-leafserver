package goclient_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3555")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{
		"Login": {
			"Account": "test",
			"Password": "123456"
		}
	}`)

	encryptData, err := RsaEncrypt(data)
	if err != nil {
		fmt.Println("加密失败:", err.Error())
		return
	}
	// len + data
	m := make([]byte, 2+len(encryptData))

	fmt.Println(encryptData)
	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(encryptData)))

	copy(m[2:], encryptData)

	// 发送消息
	_, err = conn.Write(m)
	if err != nil {
		fmt.Println("发送消息失败:", err.Error())
		return
	}

	time.Sleep(time.Second * 3)
	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes)
	if err != nil {
		fmt.Println("读取服务器消息失败", err.Error())
		return
	}
	fmt.Println("n:", n)

	fmt.Println("result:", string(bytes[3:]))

}

const PubKey = "server_pub.pem"

// 使用公钥加密
func RsaEncrypt(origData []byte) (data []byte, err error) {

	pubKey, err := ioutil.ReadFile(PubKey)
	if err != nil {
		return
	}

	//解密pem格式的公钥
	block, _ := pem.Decode(pubKey)

	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	//pub, err := x509.ParsePKCS1PublicKey(block.Bytes)

	//类型断言
	pub := pubInterface.(*rsa.PublicKey)

	data, err = rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

	return
}
