###阿里云邮件推送API
*安装：
`go get github.com/YY-tomorrow/aliyun-Mail-go`
*使用
`	profile := mail.GetProfile("cn-hangzhou", "<your accessKey>", "<your accessSecret>")
 	client := mail.GetClient(profile)
 	client.Action("SingleSendMail")
 	client.AccountName("控制台创建的发信地址")
 	client.ToAddress("收信地址")
 	client.HtmlBody("邮件正文")
 	client.Subject("邮件主题")
 	client.AddressType(1)
 	client.ReplyToAddress(true)
 	fmt.Println(client.Send())`