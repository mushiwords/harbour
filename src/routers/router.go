package routers
import (
	"github.com/astaxie/beego"
	"monitor/src/controllers"
)

func init() {
	beego.Router("/monitor", &controllers.MainController{})
	beego.Router("/monitor/captain", &controllers.CaptainHandler{})
}