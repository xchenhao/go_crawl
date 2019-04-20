package utils

import "github.com/astaxie/beego/httplib"

func Get(url string, headers map[string]string) string {
	rsp := httplib.Get(url)
	for k, v := range headers {
		rsp.Header(k, v)
	}

	htmlContent, err := rsp.String()
	if err != nil {
		panic(err)
	}
	return htmlContent
}

func InArray(item string, array []string) bool {
	for _, v := range array {
		if item == v {
			return true
		}
		continue
	}
	return false
}
