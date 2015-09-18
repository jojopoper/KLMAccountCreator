package routers

import (
	"github.com/astaxie/beego"
	"jojopoper/KLMAccountCreator/controllers"
)

func init() {
	beego.Router("/new_account", &controllers.MainController{})
}
