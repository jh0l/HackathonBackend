package controllers

import (
	"github.com/astaxie/beego"
	"hackathon/services"
)

type GetQuestionController struct {
	beego.Controller
}

type GetQuestionResponse struct {
	Uri string `json:"uri"`
	Options []string `json:"options"`
}



func (this *GetQuestionController) Get() {
	//abc := this.GetString("abc")
	//fmt.Println(abc)

	uri, options := services.GetRandomLogos()

	response := GetQuestionResponse{
		Uri: uri,
		Options: options,
	}
	this.Data["json"] = response
	this.ServeJSON()

}