package controllers

import (
	"encoding/json"
	"hackathon/services"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (this *LoginController) Login() {
	var request LoginRequest
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
		// something is wrong when parse the request
		res := LoginResponse{
			Status:  "Fail",
			Message: "invalid request: " + err.Error(),
			Token:   "",
		}
		this.Data["json"] = res
		this.ServeJSON()
		return
	}

	username := request.Username
	password := request.Password
	token, err := services.VerifyLogin(username, password)

	if err != nil {
		res := LoginResponse{
			Status:  "Fail",
			Message: "invalid request: " + err.Error(),
			Token:   "",
		}
		this.Data["json"] = res
		this.ServeJSON()
		return
	}

	res := LoginResponse{
		Status:  "Success",
		Message: "Login Successful",
		Token:   token,
	}
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *LoginController) Signup() {
	var request LoginRequest
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &request); err != nil {
		res := LoginResponse{
			Status:  "Fail",
			Message: "invalid request" + err.Error(),
			Token:   "",
		}
		this.Data["json"] = res
		this.ServeJSON()
		return
	}

	token, err := services.NewAccount(request.Username, request.Password)

	if err != nil {
		res := LoginResponse{
			Status:  "Fail",
			Message: "invalid request " + err.Error(),
			Token:   "",
		}
		this.Data["json"] = res
		this.ServeJSON()
		return
	}

	res := LoginResponse{
		Status:  "Success",
		Message: "Account Created",
		Token:   token,
	}

	this.Data["json"] = res
	this.ServeJSON()
}
