package mail

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type client struct {
	profile        *profile
	accountName    string
	addressType    string
	replyToAddress string
	subject        string
	toAddress      string
	action         string
	clickTrace     string
	fromAlias      string
	htmlBody       string
	tagName        string
	textBody       string
}

func GetClient(profile *profile) *client {
	return &client{
		profile: profile,
	}
}

func (t *client) AccountName(s string) {
	t.accountName = s
}

func (t *client) AddressType(s int) {
	t.addressType = strconv.Itoa(s)
}

func (t *client) ReplyToAddress(s bool) {
	t.replyToAddress = strconv.FormatBool(s)
}

func (t *client) Subject(s string) {
	t.subject = s
}

func (t *client) ToAddress(s string) {
	t.toAddress = s
}

func (t *client) Action(s string) {
	t.action = s
}

func (t *client) ClickTrace(s string) {
	t.clickTrace = s
}

func (t *client) FromAlias(s string) {
	t.fromAlias = s
}

func (t *client) HtmlBody(s string) {
	t.htmlBody = s
}

func (t *client) TagName(s string) {
	t.tagName = s
}

func (t *client) TextBody(s string) {
	t.textBody = s
}

var Aliyun map[string]interface{}

func (t *client) Send() map[string]interface{} {
	urldata := url.Values{}
	//获取utc时间
	utc := GetUtcTime()
	//把请求参数一一添加起来
	urldata.Add("Action", t.action)
	urldata.Add("AccountName", t.accountName)
	urldata.Add("ReplyToAddress", t.replyToAddress)
	urldata.Add("AddressType", t.addressType)
	urldata.Add("ToAddress", t.toAddress)
	urldata.Add("FromAlias", t.fromAlias)
	urldata.Add("Subject", t.subject)
	urldata.Add("ClickTrace", t.clickTrace)
	urldata.Add("HtmlBody", t.htmlBody)
	urldata.Add("TagName", t.tagName)
	urldata.Add("TextBody", t.textBody)
	urldata.Add("Format", "JSON")
	urldata.Add("Version", "2015-11-23")
	urldata.Add("SignatureMethod", "HMAC-SHA1")
	urldata.Add("SignatureNonce", utc)
	urldata.Add("SignatureVersion", "1.0")
	urldata.Add("AccessKeyId", t.profile.AccessKey)
	urldata.Add("Timestamp", utc)
	//参数encode编码,并替换参数中的特殊字符
	percent := PercentEncode(urldata.Encode())
	//获取待签名数据
	StringToSign := "GET" + "&" + url.QueryEscape("/") + "&" + url.QueryEscape(percent)
	//生成签名值
	Signature := CreateSignature(t.profile.Secret, StringToSign)
	//把签名内容带到请求参数中
	urldata.Add("Signature", Signature)

	//发起get请求
	res, err := http.Get("https://dm.aliyuncs.com/?" + PercentEncode(urldata.Encode()))
	if err != nil {
		return map[string]interface{}{
			"Error": err.Error(),
		}
	}
	//读取完成后关闭链接
	defer res.Body.Close()

	//读取返回的数据
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return map[string]interface{}{
			"Error": err.Error(),
		}
	}

	//json解析响应的数据
	json.Unmarshal(data, &Aliyun)
	return Aliyun
}
