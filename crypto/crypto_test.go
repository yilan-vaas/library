package crypto

import (
	"fmt"
	"testing"
)

var vid int64 = 10000001
var str string = "hello"
var salt string = "yilanvaas0123456"

func TestAES(t *testing.T) {
	ret, err := AesEncrypt(str, salt)
	decode, err := AesDecrypt(ret, salt)
	fmt.Println(ret, decode, err)
}

func TestHash(t *testing.T) {
	str, err := EncodeHash(vid, salt)
	id, err := DecodeHash(str, salt)
	fmt.Println(str, id, err)
}
