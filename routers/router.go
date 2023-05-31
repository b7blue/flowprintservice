package routers

import (
	"flowprintservice/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signup", &controllers.SignupController{})
	beego.Router("/sessdisplay", &controllers.SessDisplayController{})
	beego.Router("/manageprints", &controllers.ManagePrintsController{})
	beego.Router("/manageoneprint", &controllers.ManageOnePrintController{})
}
