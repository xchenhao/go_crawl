package routers

import (
	"github.com/xchenhao/go_crawl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/crawl_movie", &controllers.CrawlMovieController{}, "*:CrawlMovie")
   	beego.Router("/crawl_bbs", &controllers.CrawlBBSController{}, "*:CrawlBBS")

   	beego.Router("/test", &controllers.CrawlBBSController{}, "*:Test")
}
