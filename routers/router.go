package routers

import (
	"go_crawl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/crawl", &controllers.CrawlMovieController{}, "*:CrawlMovie")
}
