package mail

type profile struct {
	Region    string
	AccessKey string
	Secret    string
}

func GetProfile(region, accessKey, secret string) *profile {
	return &profile{
		Region:    region,
		AccessKey: accessKey,
		Secret:    secret,
	}
}
