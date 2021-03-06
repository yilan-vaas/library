package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

//errors define
var (
	pubKeyErr error = errors.New("public key error")
	priKeyErr error = errors.New("private key error")
)

//RSA Encrypt
func RSAEncrypt(data, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, pubKeyErr
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

//RSA Decrypt
func RSADecrypt(cipher, priKey []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, priKeyErr
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipher)
}

//RSA Encrypt With Base64
func RSAEncryptStr(data, pubKey string) (string, error) {
	ret, err := RSAEncrypt([]byte(data), []byte(pubKey))
	str := base64.StdEncoding.EncodeToString(ret)
	return str, err
}

//RSA Decrypt With Base64
func RSADecryptStr(str, priKey string) (string, error) {
	oriData, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	ret, err := RSADecrypt(oriData, []byte(priKey))
	data := string(ret)
	return data, err
}
