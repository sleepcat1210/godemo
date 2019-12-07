package routers

import (
	"yjg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UsersController{},"get:Get")
    beego.Router("/login", &controllers.UsersController{},"post:Login")
    beego.Router("/add/img", &controllers.UsersController{},"post:UploadImg")
}
