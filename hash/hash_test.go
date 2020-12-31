package hash

import (
	"fmt"
	"testing"
)

var str = "hello123"

func TestMD5(t *testing.T) {
	ret, err := MD5(str)
	fmt.Println(ret, err)
}

func TestSHA1(t *testing.T) {
	ret, err := SHA1(str)
	fmt.Println(ret, err)
}

func TestSHA256(t *testing.T) {
	ret, err := SHA256(str)
	fmt.Println(ret, err)
}

func TestSHA512(t *testing.T) {
	ret, err := SHA512(str)
	fmt.Println(ret, err)
}
