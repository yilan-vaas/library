package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

//MD5
func MD5(str string) (s string, err error) {
	h := md5.New()
	if _, err = h.Write([]byte(str)); err == nil {
		return hex.EncodeToString(h.Sum(nil)), nil
	}
	return
}

//SHA1
func SHA1(str string) (s string, err error) {
	h := sha1.New()
	if _, err = h.Write([]byte(str)); err == nil {
		return hex.EncodeToString(h.Sum(nil)), nil
	}
	return
}

//SHA256
func SHA256(str string) (s string, err error) {
	h := sha256.New()
	if _, err = h.Write([]byte(str)); err == nil {
		return hex.EncodeToString(h.Sum(nil)), nil
	}
	return
}

//SHA512
func SHA512(str string) (s string, err error) {
	h := sha512.New()
	if _, err = h.Write([]byte(str)); err == nil {
		return hex.EncodeToString(h.Sum(nil)), nil
	}
	return
}

//HashMac
func HMac256(str, key string) (s string, err error) {
	h := hmac.New(sha256.New, []byte(key))
	if _, err = h.Write([]byte(str)); err == nil {
		return hex.EncodeToString(h.Sum(nil)), nil
	}
	return
}
