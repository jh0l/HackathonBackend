package routers

import (
	"hackathon/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/get_question", &controllers.GetQuestionController{}, "get:Get")
	beego.Router("/verify_answer", &controllers.VerifyAnswerController{}, "post:Post")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/signup", &controllers.LoginController{}, "post:Signup")

	beego.Router("/api/charge_points", &controllers.ChargePointController{}, "get:List", "post:Add")
	beego.Router("/api/charge_points/:id([0-9]+)", &controllers.ChargePointController{}, "get:GetInfo", "post:Edit", "delete:Delete")
	beego.Router("/api/charge_points/:id([0-9]+)/sockets", &controllers.ChargePointController{}, "get:ListSockets", "post:AddSocket")
	beego.Router("/api/charge_points/:id([0-9]+)/sockets/:socketId([0-9]+)", &controllers.ChargePointController{}, "get:GetSocketInfo", "post:EditSocket", "delete:DeleteSocket")
	beego.Router("/api/charge_points/:id([0-9]+)/sockets/:socketId([0-9]+)/occupy", &controllers.ChargePointController{}, "post:OccupySocket")
	beego.Router("/api/charge_points/:id([0-9]+)/sockets/:socketId([0-9]+)/free", &controllers.ChargePointController{}, "post:FreeSocket")
	beego.Router("/api/charge_points/:id([0-9]+)/batteries", &controllers.ChargePointController{}, "get:ListBatteries", "post:AddBattery")
	beego.Router("/api/charge_points/:id([0-9]+)/batteries/:batteryId([0-9]+)", &controllers.ChargePointController{}, "get:GetBatteryInfo", "post:EditBattery", "delete:DeleteBattery")
	beego.Router("/api/charge_points/:id([0-9]+)/batteries/:batteryId([0-9]+)/take", &controllers.ChargePointController{}, "post:TakeBattery")
	beego.Router("/api/charge_points/:id([0-9]+)/batteries/:batteryId([0-9]+)/restock", &controllers.ChargePointController{}, "post:RestockBattery")
}
