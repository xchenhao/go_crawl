package routers

import (
	"github.com/xchenhao/go_crawl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/crawl_movie", &controllers.CrawlMovieController{}, "*:CrawlMovie")
}
