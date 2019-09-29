package controllers

import (
	"encoding/json"
	"hackathon/services"

	"github.com/astaxie/beego"
)

type VerifyAnswerController struct {
	beego.Controller
}

type VerifyAnswerRequest struct {
	URI       string `json:"uri"`
	Selection string `json:"selection"`
}

type VerifyAnswerResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Correct bool   `json:"correct"`
}

func (this *VerifyAnswerController) Post() {
	var request VerifyAnswerRequest
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
		// something is wrong when parse the request
		res := VerifyAnswerResponse{
			Status:  "Fail",
			Message: "invalid request: " + err.Error(),
			Correct: false,
		}
		this.Data["json"] = res
		this.ServeJSON()
		return
	}

	uri := request.URI
	selection := request.Selection
	correct := services.VerifyAnswer(uri, selection)

	res := VerifyAnswerResponse{
		Status:  "success",
		Message: "You've got an answer",
		Correct: correct,
	}
	this.Data["json"] = res
	this.ServeJSON()

}
