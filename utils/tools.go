package utils

import "github.com/astaxie/beego/httplib"

func Get(url string, ua string, cookie string) string {
	rsp := httplib.Get(url)

	rsp.Header("User-Agent", ua)
	rsp.Header("Cookie", cookie)
	htmlContent, err := rsp.String()
	if err != nil {
		panic(err)
	}
	return htmlContent
}