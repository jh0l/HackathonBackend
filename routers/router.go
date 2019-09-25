package routers

import (
	"hackathon/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/get_question", &controllers.GetQuestionController{}, "get:Get")
    beego.Router("/verify_answer", &controllers.VerifyAnswerController{}, "post:Post")
}

