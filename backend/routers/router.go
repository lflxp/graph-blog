package routers

import (
	"github.com/lflxp/graph-blog/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
