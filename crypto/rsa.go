package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 加密
func RSAEncrypt(data, pubKey string) (str string, err error) {
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return "", errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	pub := pubInterface.(*rsa.PublicKey)
	ret, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(data))
	if err == nil {
		str = base64.StdEncoding.EncodeToString(ret)
	}
	return
}

// 解密
func RASDecrypt(str, priKey string) (data string, err error) {
	block, _ := pem.Decode([]byte(priKey))
	if block == nil {
		return "", errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	oriData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}

	ret, err := rsa.DecryptPKCS1v15(rand.Reader, priv, oriData)
	data = string(ret)
	return
}
