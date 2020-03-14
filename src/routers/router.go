package routers
import (
	"harbour/src/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/captain", &controllers.CaptainHandler{})
}