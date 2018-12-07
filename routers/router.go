package routers

import (
	"github.com/astaxie/beego"
	"github.com/justforlxz/mkacg/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/post", &controllers.PostController{})
	beego.Router("/categories", &controllers.CategoriesController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
}
