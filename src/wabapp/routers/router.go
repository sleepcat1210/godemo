package routers

import (
	"wabapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	
	beego.Router("/user", &controllers.UserController{},"get:ShowUser")
	beego.Router("/user/add", &controllers.UserController{},"post:AddUser")
	beego.Router("/user/edit", &controllers.UserController{},"get:Edit;post:EditInfo")
	beego.Router("/user/img", &controllers.UserController{},"get:GetImg")
	beego.Router("/info", &controllers.UserController{},"get:UserInfo")
	beego.Router("/reg", &controllers.UserController{},"get:Register")
	beego.Router("/del", &controllers.UserController{},"get:Del")
	beego.Router("/address", &controllers.AddressController{},"get:GetAddress")
	beego.Router("/address/add", &controllers.AddressController{},"get:Add;post:Addr")

}
