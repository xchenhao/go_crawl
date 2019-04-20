package controllers

import (
	"github.com/xchenhao/go_crawl/models"
	"github.com/astaxie/beego"
	"github.com/xchenhao/go_crawl/utils"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}

/**
 目前这个爬虫只能爬取静态数据 对于像京东的部分动态数据 无法爬取
 对于动态数据 可以采用 一个组件 phantomjs
*/
func (c *CrawlMovieController) CrawlMovie() {
	//爬虫入口url
	sUrl := "https://movie.douban.com/subject/6874741/?from=subject-page"
	models.PutinQueue(sUrl, models.KEY_DOUBAN_MOVIE_URL_QUEUE)

	for {
		// 如果 url 队列为空 则退出循环
		if 0 == models.GetQueueLength(models.KEY_DOUBAN_MOVIE_URL_QUEUE) {
			break
		}

		sUrl = models.PopfromQueue(models.KEY_DOUBAN_MOVIE_URL_QUEUE)
		// 判断 sUrl 是否被访问过
		if models.IsVisit(sUrl, models.KEY_DOUBAN_MOVIE_URL_VISIT_SET) {
			continue
		}

		// 设置 User-agent 以及 Cookie 是防止豆瓣网的 403
		var header = make(map[string]string, 2)
		header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0"
		header["Cookie"] = `bid=gFP9qSgGTfA; __utma=30149280.1124851270.1482153600.1483055851.1483064193.8; __utmz=30149280.1482971588.4.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; ll="118221"; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1483064193%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=5afcf5e5496eab22.1482413017.7.1483066280.1483057909.; __utma=223695111.1636117731.1482413017.1483055857.1483064193.7; __utmz=223695111.1483055857.6.5.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=BDC2DBEDF8958EC838F9D9394CC5D9A0|2cc6ef7952be8c2d5408cb7c8cce2684; ap=1; viewed="1006073"; gr_user_id=e5c932fc-2af6-4861-8a4f-5d696f34570b; __utmc=30149280; __utmc=223695111; _pk_ses.100001.4cf6=*; __utmb=30149280.0.10.1483064193; __utmb=223695111.0.10.1483064193`


		sMovieHtml := utils.Get(sUrl, header)
		// 添加电影记录
		models.AddMovie(sMovieHtml)
		// sUrl 记录到访问过的集体当中
		models.AddToSet(sUrl, models.KEY_DOUBAN_MOVIE_URL_VISIT_SET)

		// 提取该页面的所有连接
		urls := models.GetMovieUrls(sMovieHtml)
		for _, url := range urls {
			models.PutinQueue(url, models.KEY_DOUBAN_MOVIE_URL_QUEUE)
			c.Ctx.WriteString("<br>" + url + "</br>")
		}

		time.Sleep(time.Second)
	}

	c.Ctx.WriteString("end of crawl!")
}
