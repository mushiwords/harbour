package routers
import (
	"github.com/astaxie/beego"
	"monitor/src/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/captain", &controllers.CaptainHandler{})
}