package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xchenhao/go_crawl/models"
	"github.com/xchenhao/go_crawl/utils"
	"time"
)

type CrawlBBSController struct {
	beego.Controller
}

var (
	bbs_fid = 80

	start_page = 1
	end_page = 9
	urlTpl = "https://www.mukedaba.com/forum.php?mod=forumdisplay&fid=%d&page=%d"
)

func (c *CrawlBBSController) CrawlBBS() {

	var header = make(map[string]string, 10)
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0"
	header["Cookie"] = `_uab_collina=155098388201235580002924; a4Ib_2132_connect_is_bind=1; a4Ib_2132_nofavfid=1; a4Ib_2132_saltkey=sC0Ucj9Z; a4Ib_2132_lastvisit=1553733326; a4Ib_2132_auth=affcpwsbvjt17lCv%2B0nMFvP4PBD5eDtTmbZ6VOpDG2EL2kGT7BdVQ%2Fm1RviwLZSxIwVsTI9jp65jdn%2FS4AbRHtnTEzo; a4Ib_2132_lastcheckfeed=115507%7C1553736958; a4Ib_2132_sid=Ukqnqb; a4Ib_2132_pc_size_c=0; a4Ib_2132_ulastactivity=3213bOJHde7%2BdaHvno6Z6imV19G9NyFcOUaGtGfI19Ht3ow1NxQv; a4Ib_2132_noticeTitle=1; a4Ib_2132_visitedfid=80D76D79D82D78D96D84D92; a4Ib_2132_home_diymode=1; a4Ib_2132_lastact=1555734418%09forum.php%09forumdisplay; a4Ib_2132_st_t=115507%7C1555734418%7C4e3e61acc8a22275943f306118b1210a; a4Ib_2132_forum_lastvisit=D_76_1555734335D_80_1555734418`
	header["Host"] = "www.mukedaba.com"
	header["DNT"] = "1"
	header["Referer"] = "https://www.mukedaba.com/forum-" + string(bbs_fid) + "-1.html"
	header["X-Requested-With"] = "XMLHttpRequest"
	header["Accept"] = "*/*"
	header["Connection"] = "keep-alive"

	// 获取版块的各个帖子 URL
	c.Ctx.WriteString("crawl_urls\n\r")
	for i := start_page; i <= end_page; i++ {
		url := fmt.Sprintf(urlTpl, bbs_fid, i)
		htmlContent := 	utils.Get(url, header)

		urlList := models.GetUrls(htmlContent)
		for k := 0; k < len(urlList); k++ {
			models.AddToSet(urlList[k], models.KEY_MKDB_BBS_URL_QUEUE)
		}
		c.Ctx.WriteString(url + "\n\r")

		time.Sleep(time.Second * 2)
	}
	c.Ctx.WriteString("crawl_content\n\r")

	var skip_array = []string {
		"https://www.mukedaba.com/thread-23207-1-1.html",
		"https://www.mukedaba.com/thread-2177-1-1.html",
		"https://www.mukedaba.com/thread-1653-1-1.html",
	}
	// 爬取帖子内容
	for {
		if 0 == models.GetSetLenth(models.KEY_MKDB_BBS_URL_QUEUE) {
			break
		}
		cUrl := models.PopFromSet(models.KEY_MKDB_BBS_URL_QUEUE)
		if "" == cUrl || utils.InArray(cUrl, skip_array) ||  models.IsVisit(cUrl, models.KEY_MKDB_BBS_URL_VISIT_SET) {
			continue
		}
		c.Ctx.WriteString(cUrl + "\n\r")

		postContent := utils.Get(cUrl, header)
		models.AddPost(postContent)

		models.AddToSet(cUrl, models.KEY_MKDB_BBS_URL_VISIT_SET)
		time.Sleep(time.Second * 2)
	}

	c.Ctx.WriteString("end of crawl!")
}

func (c *CrawlBBSController) Test() {
	//data, err := ioutil.ReadFile("post.html")
	//if  nil != err {
	//	panic(err)
	//}
	var header = make(map[string]string, 10)
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0"
	header["Cookie"] = `_uab_collina=155098388201235580002924; a4Ib_2132_connect_is_bind=1; a4Ib_2132_nofavfid=1; a4Ib_2132_saltkey=sC0Ucj9Z; a4Ib_2132_lastvisit=1553733326; a4Ib_2132_auth=affcpwsbvjt17lCv%2B0nMFvP4PBD5eDtTmbZ6VOpDG2EL2kGT7BdVQ%2Fm1RviwLZSxIwVsTI9jp65jdn%2FS4AbRHtnTEzo; a4Ib_2132_lastcheckfeed=115507%7C1553736958; a4Ib_2132_sid=Ukqnqb; a4Ib_2132_pc_size_c=0; a4Ib_2132_ulastactivity=3213bOJHde7%2BdaHvno6Z6imV19G9NyFcOUaGtGfI19Ht3ow1NxQv; a4Ib_2132_noticeTitle=1; a4Ib_2132_visitedfid=80D76D79D82D78D96D84D92; a4Ib_2132_home_diymode=1; a4Ib_2132_lastact=1555734418%09forum.php%09forumdisplay; a4Ib_2132_st_t=115507%7C1555734418%7C4e3e61acc8a22275943f306118b1210a; a4Ib_2132_forum_lastvisit=D_76_1555734335D_80_1555734418`
	header["Host"] = "www.mukedaba.com"
	header["DNT"] = "1"
	header["Referer"] = "https://www.mukedaba.com/forum-" + string(bbs_fid) + "-1.html"
	header["X-Requested-With"] = "XMLHttpRequest"
	header["Accept"] = "*/*"
	header["Connection"] = "keep-alive"

	cUrl := "https://www.mukedaba.com/thread-3266-1-6.html"
	data := utils.Get(cUrl, header)

	content := models.GetPostContent(string(data))

	models.AddPost(data)
	c.Ctx.WriteString(content)
}
