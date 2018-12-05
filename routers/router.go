package routers

import (
	"github.com/justforlxz/mkacg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
