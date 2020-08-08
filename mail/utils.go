package mail

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

func PercentEncode(str string) string {
	//替换字符串
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func CreateSignature(secret, StringToSign string) string {
	//创建签名
	key := []byte(secret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(StringToSign))
	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return s
}

func GetUtcTime() string {
	//获取utc时间
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	s := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", year, mon, day, hour, min, sec)
	return s
}
