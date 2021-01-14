package crypto

import (
	"fmt"
	"strings"

	"github.com/speps/go-hashids"
)

//最短长度
var minLen int = 12

func EncodeHash(id int64, salt string) (string, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLen
	hash, err := hashids.NewWithData(hd)
	if err != nil {
		return "", fmt.Errorf("encode hash err: %s", err)
	}

	arr := []int{int(id)}
	str, err := hash.Encode(arr)
	if err != nil {
		return "", fmt.Errorf("encode hash err: %s", err)
	}
	return str, nil
}

func DecodeHash(str, salt string) (int64, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0, fmt.Errorf("decode hash err: empty")
	}

	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLen
	hash, err := hashids.NewWithData(hd)
	if err != nil {
		return 0, fmt.Errorf("decode hash err: %s", err)
	}
	arr, err := hash.DecodeWithError(str)
	if err != nil {
		return 0, fmt.Errorf("decode hash err: %s", err)
	}
	return int64(arr[0]), nil
}
