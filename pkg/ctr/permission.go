package ctr

var accessWhiteList []string

func InitAccessWhiteList(whiteList []string) {
	accessWhiteList = whiteList
}

// 获取访问白名单
func AccessWhiteList() []string {
	return accessWhiteList
}
