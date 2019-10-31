package cryptoutils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func GetGuid() string {

	f, err := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)

	if err == nil {
		b := make([]byte, 16)
		f.Read(b)
		f.Close()
		uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
		return string(uuid)

	}

	s := NewV4()
	return fmt.Sprintf("%s", s)

}

func Get16Md5Encode(src string) string {
	runes := []rune(Md5Encode(src))
	return string(runes[7:23])
}

func Md5Encode(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
