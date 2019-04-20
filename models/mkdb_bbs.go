package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

const (
	KEY_MKDB_BBS_URL_QUEUE = "mkdb_url_queue"
	KEY_MKDB_BBS_URL_VISIT_SET = "mkdb_url_visited_set"
)

var (
	mk_dbConnection orm.Ormer = GetConnection(new(MkPosts), "default")
)
// mk_posts
type MkPosts struct {
	Id int64
	Title string
	Content string
}

func GetUrls(html string) []string {
	var urlSet []string

	if "" == html {
		return urlSet
	}

	// https://www.mukedaba.com/thread-3200-1-6.html
	reg := regexp.MustCompile(`.*?"(http.+?www\.mukedaba\.com/thread.*?html).*?".*?`)
	result := reg.FindAllStringSubmatch(html, -1)

	for _,v := range result{
		urlSet = append(urlSet, v[1])
	}

	return urlSet
}

func AddPost(html string) (int64, error) {
	postTitle := getPostTitle(html)
	if "" == postTitle {
		return 0, nil
	}

	var model MkPosts
	content := GetPostContent(html)
	fmt.Println(content)
	model.Title          = postTitle
	model.Content        = content
	model.Id = 0

	id, err := mk_dbConnection.Insert(&model)
	return id, err
}

func getPostTitle(html string) string {
	if html == ""{
		return ""
	}

	reg := regexp.MustCompile(`.*?<span id="thread_subject".*?>(.+?)</span>.*?`)
	result := reg.FindAllStringSubmatch(html, -1)

	if 0 == len(result) {
		return ""
	}

	return string(result[0][1])
}

func GetPostContent(html string) string {
	if html == ""{
		return ""
	}

	reg := regexp.MustCompile(`(?s:.+?<table cellspacing="0" cellpadding="0".+?postmessage.+?>(.+?class="showhide".+?)</table>.+?class="tshare cl".+?)`)
	result := reg.FindAllStringSubmatch(html, -1)

	if 0 == len(result) {
		return ""
	}

	return string(result[0][1])
}
