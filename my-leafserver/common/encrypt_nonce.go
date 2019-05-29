package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

//const (
//	PriKey = "cert/private.pem"
//	PubKey = "cert/public.pem"
//)
const publicKey  = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1PvP0k0bP9oJd2k9lUVt
u7YH7q0jmgcUPqL4jOQen6D37mFOKF4fhcqSdt65CbSgdjGP4Kln6Mvy4BHAR54i
a4fXmbxXb9pc4mhnTFGoi5eXGgcs/VfR2M5DE2Ik8Hy7XU3uhf74iVCW9MGU4LJZ
ehAa478SPHkaEnf2us0AFlbyp2IXOp11ZpIvUEsr7JC1E1wsnABVcORjSlzCE2Sz
MEBKjLBYOTDuBXRnpLCkgl1OPJ215t5FVo0tAc3LGzjNAR0xrhEAI0cKGofZaXUF
7H7ifO4Hmk/hAVv45+f3gEcqlJ5uKMoB34BH8ejTwx30SupHFP06E4c5uNGjqZ4U
ZwIDAQAB
-----END RSA PUBLIC KEY-----`

const privateKey  =`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA1PvP0k0bP9oJd2k9lUVtu7YH7q0jmgcUPqL4jOQen6D37mFO
KF4fhcqSdt65CbSgdjGP4Kln6Mvy4BHAR54ia4fXmbxXb9pc4mhnTFGoi5eXGgcs
/VfR2M5DE2Ik8Hy7XU3uhf74iVCW9MGU4LJZehAa478SPHkaEnf2us0AFlbyp2IX
Op11ZpIvUEsr7JC1E1wsnABVcORjSlzCE2SzMEBKjLBYOTDuBXRnpLCkgl1OPJ21
5t5FVo0tAc3LGzjNAR0xrhEAI0cKGofZaXUF7H7ifO4Hmk/hAVv45+f3gEcqlJ5u
KMoB34BH8ejTwx30SupHFP06E4c5uNGjqZ4UZwIDAQABAoIBAFjPqBlJjdNSWCAz
Ajr8eIWNokEkcXQI/6gezQXzGdH2jWwZpz2uAfcIQYrP1Nj3OPIQOK619V5drSlC
PaufhTrRqlWw88TwiUrfvjNU2bNbwuZXl6sMs7R0TgQHKOAfIwIFPTEi4QBhWpE1
J5+Kv0Tm8k+FnVUaKoUJARlj4mVtmW2IRLmHQFNbOObIc6spOCDw9gODx0NBO9bC
4Y0wXJBWchto36ez6FRyUtXV/9CtgQei9mpcmrhMh6s8kY7I5MJrve4tzobNeAeA
YrFZRqXO6Y9zTlJnyKKmS305enCTaBLTSGI6Rn0yzqmwrB+QK6EUpbY2Yt0CVAAS
WVaXfgECgYEA4ny4jVX41mhoeVGKsI8SR0aSrPoOMUz+VmGCqghyYn6+9Miy6Gui
QTdkKDg+Gb/DFcCRTbroKHTAdxhCKTywqyenp0G+yEkhL3oRFkkmpEw5zmh8tI5K
eUpOHdUq1vOCRJ68naSULWbK0haIiDX4KpKt3MqVFSv2f6ZOWDgLhgUCgYEA8Lyh
35v8z049+uwnWoHLcBpXX12oTD4yzIjwWT29KlEaMVOf9L5iwk9QB5c7LqtlS8V0
JOY8z1i9H4HJagGoBio9x+6b/c3lHZPx3taZ/0+7Cs1tcsTFAH91JHucOyKFt3A6
48y6i2iVkvePdCrb8NPP1hrH1ZWuPovbyY+w8HsCgYEAi1nXp0ZlU25sEXEFVzrA
MVCPwJQeFWFc/8MBRSLrVgTL8wplYbGP5HZzTnfEZc1h98lC3cOJQhMLtHsdggfQ
X46HjdZazxqHq5F8X+zkNKSd0IzjZahC1DvOsnZM5HJxDU1pmckATqr3UaoBmWUz
auZSyQ1wCtXA1at6ercJCA0CgYB4LAG2SPzTU2B3QKmIcaBnTo7tCSi5HpdnKLiq
I7qpRCEKHI+NsfhEvSjbETA25NJFF8UZomEatFZ8QuKTa04//ZgnlMID2WMU98RW
k9P98gQRqWiIURdyXy3Pz3C2yE3tuzV6f7ljXArGeZP/zmFKf8GRRH8a7IQ1rsVB
gdw2vwKBgBQBFghkJLO+yQvrR29nIiplQ97X+rjTZWZcXbGURtRXJd+A8XAvCccr
mKUKQH71R1wnOXAVeoJV4Y/pMhjbkcHZQs7mmbtdjQLaQuG1YcF8z4PLna8TEGMD
dZVDJB1mQ0ghMod0PMyDcM/+vqFIUsMdaLO8KZvkqLaF5cjzDABK
-----END RSA PRIVATE KEY-----`




var PubPerm []byte
var PriPerm []byte

func init() {
	//var err error
	//PriPerm, err = ioutil.ReadFile(PriKey)
	//if err != nil {
	//	logger.Error("init rsa priKey fail!")
	//	panic(err.Error())
	//}
	//
	//PubPerm, err = ioutil.ReadFile(PubKey)
	//if err != nil {
	//	logger.Error("init rsa pubKey fail!")
	//	panic(err.Error())
	//}
	PubPerm =[]byte(publicKey)
	PriPerm =[]byte(privateKey)

}

// 使用公钥加密
func RsaEncrypt(origData ,perm []byte) (data []byte, err error) {
	//pubKey, err := ioutil.ReadFile(PubKey)
	//
	////pubKey, err := ioutil.ReadFile(PubKey)
	//if err != nil {
	//	return
	//}

	//解密pem格式的公钥
	block, _ := pem.Decode(perm)

	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	//pub, err := x509.ParsePKCS1PublicKey(block.Bytes)

	//类型断言
	pub := pubInterface.(*rsa.PublicKey)

	data, err = rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

	return
}

// 使用私钥解密
func RsaDecrypt(cipherText ,perm []byte) (data []byte, err error) {
	//priKey, err := ioutil.ReadFile(PriKey)
	//if err != nil {
	//	return
	//}

	//解密
	block, _ := pem.Decode(perm)
	if block == nil {
		return
	}

	//解析PKCS1格式的私钥
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	data, err = rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
	return
}
